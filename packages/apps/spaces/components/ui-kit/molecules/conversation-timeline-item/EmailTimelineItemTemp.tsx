import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { toast } from 'react-toastify';
import { EmailTimelineItem } from '../email-timeline-item';
import { TimelineItem } from '../../atoms/timeline-item';
import { Skeleton } from '../../atoms/skeleton';
import { useRecoilState, useRecoilValue, useSetRecoilState } from 'recoil';
import {
  editorEmail,
  editorMode,
  EditorMode,
  userData,
} from '../../../../state';
import { ConversationItem, FeedPostRequest, Props } from './types';

export const EmailTimelineItemTemp: React.FC<Props> = ({
  feedId,
  first,
  feedInitiator,
}) => {
  const setEditorMode = useSetRecoilState(editorMode);
  const [emailEditorData, setEmailEditorData] = useRecoilState(editorEmail);
  const loggedInUserData = useRecoilValue(userData);
  const [messages, setMessages] = useState([] as ConversationItem[]);

  const [loadingMessages, setLoadingMessages] = useState(false);

  useEffect(() => {
    return () => {
      setEmailEditorData({
        //@ts-expect-error fixme later
        handleSubmit: () => null,
        to: [],
        subject: '',
        respondTo: '',
      });
      setEditorMode({
        mode: EditorMode.Note,
        submitButtonLabel: 'Log into timeline',
      });
    };
  }, []);

  useEffect(() => {
    setLoadingMessages(true);
    axios
      .get(`/oasis-api/feed/${feedId}/item`)
      .then((res) => {
        setMessages(res.data ?? []);
        setLoadingMessages(false);
      })
      .catch(() => {
        setLoadingMessages(false);
        toast.error('Something went wrong while loading feed item', {
          toastId: `conversation-timeline-item-feed-${feedId}`,
        });
      });
  }, []);

  const handleSendMessage = (
    text: string,
    onSuccess: () => void,
    destination = [],
    replyTo: null | string,
  ) => {
    if (!text) return;
    const message: FeedPostRequest = {
      channel: 'EMAIL',
      username: loggedInUserData.identity,
      message: text,
      direction: 'OUTBOUND',
      destination: destination,
    };
    console.log('🏷️ ----- replyTo: ', replyTo);
    if (replyTo) {
      message.replyTo = replyTo;
    }

    axios
      .post(`/oasis-api/feed/${feedId}/item`, message)
      .then((res) => {
        console.log(res);
        if (res.data) {
          setMessages((messageList: any) => [...messageList, res.data]);
          onSuccess();
          setEditorMode({
            submitButtonLabel: 'Log into timeline',
            mode: EditorMode.Note,
          });
          setEmailEditorData({ ...emailEditorData, to: [], subject: '' });
          toast.success('Email sent!');
        }
      })
      .catch(() => {
        toast.error('Something went wrong while sending email');
      });
  };

  //when a new message appears, scroll to the end of container
  useEffect(() => {
    if (messages) {
      setLoadingMessages(false);
    }
  }, [messages]);

  const timeFromLastTimestamp = new Date(1970, 0, 1).setSeconds(
    feedInitiator.lastTimestamp?.seconds,
  );

  const getSortedItems = (data: Array<any>): Array<ConversationItem> => {
    return data.sort((a, b) => {
      const date1 =
        new Date(1970, 0, 1).setSeconds(a?.time?.seconds) ||
        timeFromLastTimestamp;
      const date2 =
        new Date(1970, 0, 1).setSeconds(b?.time?.seconds) ||
        timeFromLastTimestamp;
      return date2 - date1;
    });
  };
  return (
    <div className='flex flex-column h-full w-full'>
      <div className='flex-grow-1 w-full'>
        {loadingMessages && (
          <div className='flex flex-column mb-2'>
            <div className='mb-2 flex justify-content-end'>
              <Skeleton height='40px' width='50%' />
            </div>
            <div className='mb-2 flex justify-content-start'>
              <Skeleton height='50px' width='40%' />
            </div>
            <div className='flex justify-content-end mb-2'>
              <Skeleton height='45px' width='50%' />
            </div>
            <div className='flex justify-content-start'>
              <Skeleton height='40px' width='45%' />
            </div>
          </div>
        )}

        <div className='flex flex-column w-full'>
          {
            // email
            !loadingMessages &&
              getSortedItems(messages)
                .filter((msg: ConversationItem) => msg.type === 1)
                .map((msg: ConversationItem, index: number) => {
                  const emailData = JSON.parse(msg.content);

                  if (!emailData.html) {
                    return;
                  }

                  const date =
                    new Date(1970, 0, 1).setSeconds(msg?.time?.seconds) ||
                    timeFromLastTimestamp;
                  const fl =
                    first && (index === 0 || index === messages.length - 1);
                  return (
                    <TimelineItem
                      first={fl}
                      createdAt={date}
                      //@ts-expect-error fixme later
                      key={msg?.messageId?.conversationEventId}
                    >
                      <EmailTimelineItem
                        isToDeprecate
                        content={emailData.html}
                        contentType={'text/html'}
                        sentBy={emailData.from}
                        sentTo={emailData.to}
                        deprecatedCC={emailData.cc?.join(', ')}
                        deprecatedBCC={emailData.bcc?.join(', ')}
                        interactionSession={{ name: emailData.subject }}
                      />
                    </TimelineItem>
                  );
                })
          }
        </div>
      </div>
    </div>
  );
};