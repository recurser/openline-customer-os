package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	c "github.com/openline-ai/openline-customer-os/packages/server/comms-api/config"
	"github.com/openline-ai/openline-customer-os/packages/server/comms-api/model"
	"github.com/openline-ai/openline-customer-os/packages/server/comms-api/routes/ContactHub"
	s "github.com/openline-ai/openline-customer-os/packages/server/comms-api/service"
	cosModel "github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"net/http"
	"time"
)

func callEventPartyToEventParticipantInput(party *model.CallEventParty) cosModel.InteractionEventParticipantInput {
	var participantInput cosModel.InteractionEventParticipantInput
	if party.Mailto != nil {
		participantInput = cosModel.InteractionEventParticipantInput{
			Email: party.Mailto,
		}
	} else if party.Tel != nil {
		participantInput = cosModel.InteractionEventParticipantInput{
			PhoneNumber: party.Tel,
		}
	}

	return participantInput
}

func callEventPartyToSessionParticipantInput(party *model.CallEventParty) cosModel.InteractionSessionParticipantInput {
	var participantInput cosModel.InteractionSessionParticipantInput
	if party.Mailto != nil {
		participantInput = cosModel.InteractionSessionParticipantInput{
			Email: party.Mailto,
		}
	} else if party.Tel != nil {
		participantInput = cosModel.InteractionSessionParticipantInput{
			PhoneNumber: party.Tel,
		}
	}

	return participantInput
}

func callEventGetOrCreateSession(threadId string, name string, tenant string, attendants []cosModel.InteractionSessionParticipantInput, cosService s.CustomerOSService) (*string, error) {
	var err error

	sessionId, err := cosService.GetInteractionSession(&threadId, &tenant, nil)
	if err != nil {
		se, _ := status.FromError(err)
		log.Printf("failed retriving interaction session: status=%s message=%s", se.Code(), se.Message())
	} else {
		return sessionId, nil
	}

	if sessionId == nil {
		sessionChannel := "VOICE"
		sessionAppSource := "COMMS_API"
		sessionStatus := "ACTIVE"
		sessionType := "CALL"
		sessionOpts := []s.SessionOption{
			s.WithSessionIdentifier(&threadId),
			s.WithSessionChannel(&sessionChannel),
			s.WithSessionName(&name),
			s.WithSessionAppSource(&sessionAppSource),
			s.WithSessionStatus(&sessionStatus),
			s.WithSessionTenant(&tenant),
			s.WithSessionAttendedBy(attendants),
			s.WithSessionType(&sessionType),
		}
		sessionId, err = cosService.CreateInteractionSession(sessionOpts...)

		if err != nil {
			se, _ := status.FromError(err)
			log.Printf("failed creating interaction session: status=%s message=%s", se.Code(), se.Message())
			return nil, fmt.Errorf("callEventGetOrCreateSession: failed creating interaction session: %v", err)
		}
		log.Printf("interaction session created: %s", *sessionId)
	}

	return sessionId, nil
}

func getCallEventContactWithIndex(req *model.CallEvent) (string, int) {
	if req.From != nil && req.From.Tel != nil {
		return *req.From.Tel, 0
	} else if req.To != nil && req.To.Tel != nil {
		return *req.To.Tel, 1
	}
	return "", 0
}

type callProgressEventInfo struct {
	sessionId *string
	tenant    *string
	sentBy    []cosModel.InteractionEventParticipantInput
	sentTo    []cosModel.InteractionEventParticipantInput
	eventType string
	eventData string
	eventTime time.Time
}

type OpenlineCallProgressData struct {
	Version      string     `json:"version,default=1.0"`
	StartTime    *time.Time `json:"start_time,omitempty"`
	AnsweredTime *time.Time `json:"answered_time,omitempty"`
	EndTime      *time.Time `json:"end_time,omitempty"`
	Duration     *int64     `json:"duration,omitempty"`
}

func submitCallProgressEvent(event callProgressEventInfo, cosService s.CustomerOSService) (string, error) {

	channel := "VOICE"
	appSource := "COMMS_API"
	mimeTime := "application/x-openline-call-progress"
	utcTime := event.eventTime.UTC()
	eventOpts := []s.EventOption{
		s.WithTenant(event.tenant),
		s.WithContent(&event.eventData),
		s.WithContentType(&mimeTime),
		s.WithSentBy(event.sentBy),
		s.WithSentTo(event.sentTo),
		s.WithAppSource(&appSource),
		s.WithCreatedAt(&utcTime),
		s.WithEventType(&event.eventType),
	}

	eventOpts = append(eventOpts, s.WithSessionId(event.sessionId))
	eventOpts = append(eventOpts, s.WithChannel(&channel))

	response, err := cosService.CreateInteractionEvent(eventOpts...)
	if err != nil {
		return "", fmt.Errorf("submitCallProgressEvent: failed creating interaction event: %v", err)
	}
	return response.InteractionEventCreate.Id, nil
}

func addCallEventRoutes(conf *c.Config, rg *gin.RouterGroup, cosService s.CustomerOSService, hub *ContactHub.ContactHub, redisService s.RedisService) {
	rg.POST("/call_progress", func(ctx *gin.Context) {
		var req model.CallEvent
		bodyBytes, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"result": fmt.Sprintf("unable to get http body: %v", err.Error()),
			})
			return
		}
		err = json.Unmarshal(bodyBytes, &req)
		if err != nil {
			log.Printf("unable to parse json: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"result": fmt.Sprintf("unable to parse json: %v", err.Error()),
			})
			return
		}

		isActive, tenant := redisService.GetKeyInfo(ctx, "tenantKey", ctx.Request.Header.Get("X-API-KEY"))

		if !isActive {
			ctx.JSON(http.StatusForbidden, gin.H{"result": "Invalid API Key"})
			return
		}
		threadId := req.CorrelationId

		contact, index := getCallEventContactWithIndex(&req)
		subject := ""
		if index == 0 {
			subject = fmt.Sprintf("Incoming call from %s", contact)
		} else {
			subject = fmt.Sprintf("Outgoing call to %s", contact)
		}

		var sessionParticipants []cosModel.InteractionSessionParticipantInput

		sessionParticipants = append(sessionParticipants, callEventPartyToSessionParticipantInput(req.From))
		sessionParticipants = append(sessionParticipants, callEventPartyToSessionParticipantInput(req.To))

		sessionId, err := callEventGetOrCreateSession(threadId, subject, *tenant, sessionParticipants, cosService)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"result": fmt.Sprintf("Unable to create InteractionSession! reasion: %v", err),
			})
			return
		}

		var ids []string

		switch req.Event {
		case "CALL_START":
			var callStartEvent model.CallEventStart
			if err := json.Unmarshal(bodyBytes, &callStartEvent); err != nil {
				log.Printf("unable to parse json: %v", err.Error())
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"result": fmt.Sprintf("unable to parse json: %v", err.Error()),
				})
				return
			}
			eventData := OpenlineCallProgressData{
				StartTime: &callStartEvent.StartTime,
			}

			eventDataBytes, err := json.Marshal(eventData)
			if err != nil {
				log.Printf("unable to marshal json: %v", err.Error())
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"result": fmt.Sprintf("unable to marshal json: %v", err.Error()),
				})
				return
			}

			eventInfo := callProgressEventInfo{
				sessionId: sessionId,
				tenant:    tenant,
				sentBy:    []cosModel.InteractionEventParticipantInput{callEventPartyToEventParticipantInput(req.From)},
				sentTo:    []cosModel.InteractionEventParticipantInput{callEventPartyToEventParticipantInput(req.To)},
				eventType: "CALL_START",
				eventData: string(eventDataBytes),
				eventTime: callStartEvent.StartTime,
			}
			id, err := submitCallProgressEvent(eventInfo, cosService)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"result": fmt.Sprintf("Unable to submit call progress event! reasion: %v", err),
				})
				return
			}
			ids = append(ids, id)
		case "CALL_ANSWERED":
			var callAnsweredEvent model.CallEventAnswered
			if err := json.Unmarshal(bodyBytes, &callAnsweredEvent); err != nil {
				log.Printf("unable to parse json: %v", err.Error())
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"result": fmt.Sprintf("unable to parse json: %v", err.Error()),
				})
				return
			}
			eventData := OpenlineCallProgressData{
				StartTime:    &callAnsweredEvent.StartTime,
				AnsweredTime: &callAnsweredEvent.AnsweredTime,
			}
			eventDataBytes, err := json.Marshal(eventData)
			if err != nil {
				log.Printf("unable to marshal json: %v", err.Error())
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"result": fmt.Sprintf("unable to marshal json: %v", err.Error()),
				})
				return
			}
			eventInfo := callProgressEventInfo{
				sessionId: sessionId,
				tenant:    tenant,
				sentBy:    []cosModel.InteractionEventParticipantInput{callEventPartyToEventParticipantInput(req.To)},
				sentTo:    []cosModel.InteractionEventParticipantInput{callEventPartyToEventParticipantInput(req.From)},
				eventType: "CALL_ANSWERED",
				eventData: string(eventDataBytes),
				eventTime: callAnsweredEvent.AnsweredTime,
			}
			id, err := submitCallProgressEvent(eventInfo, cosService)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"result": fmt.Sprintf("Unable to submit call progress event! reasion: %v", err),
				})
				return
			}
			ids = append(ids, id)
		case "CALL_END":
			var callEndEvent model.CallEventEnd
			if err := json.Unmarshal(bodyBytes, &callEndEvent); err != nil {
				log.Printf("unable to parse json: %v", err.Error())
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"result": fmt.Sprintf("unable to parse json: %v", err.Error()),
				})
				return
			}
			eventData := OpenlineCallProgressData{
				StartTime:    callEndEvent.StartTime,
				AnsweredTime: callEndEvent.AnsweredTime,
				EndTime:      &callEndEvent.EndTime,
				Duration:     &callEndEvent.Duration,
			}
			eventDataBytes, err := json.Marshal(eventData)
			if err != nil {
				log.Printf("unable to marshal json: %v", err.Error())
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"result": fmt.Sprintf("unable to marshal json: %v", err.Error()),
				})
				return
			}
			eventInfo := callProgressEventInfo{
				sessionId: sessionId,
				tenant:    tenant,
				eventType: "CALL_END",
				eventData: string(eventDataBytes),
				eventTime: callEndEvent.EndTime,
			}
			if callEndEvent.FromCaller {
				eventInfo.sentBy = []cosModel.InteractionEventParticipantInput{callEventPartyToEventParticipantInput(req.From)}
				eventInfo.sentTo = []cosModel.InteractionEventParticipantInput{callEventPartyToEventParticipantInput(req.To)}
			} else {
				eventInfo.sentBy = []cosModel.InteractionEventParticipantInput{callEventPartyToEventParticipantInput(req.To)}
				eventInfo.sentTo = []cosModel.InteractionEventParticipantInput{callEventPartyToEventParticipantInput(req.From)}
			}
			id, err := submitCallProgressEvent(eventInfo, cosService)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"result": fmt.Sprintf("Unable to submit call progress event! reasion: %v", err),
				})
				return
			}
			ids = append(ids, id)
		}

		log.Printf("message item created with ids: %v", ids)

		ctx.JSON(http.StatusOK, gin.H{
			"result": "success",
			"ids":    ids,
		})
	})
}