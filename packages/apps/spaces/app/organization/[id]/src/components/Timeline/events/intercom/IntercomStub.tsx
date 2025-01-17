'use client';
import React, { FC } from 'react';
import {
  ContactParticipant,
  JobRoleParticipant,
  UserParticipant,
} from '@graphql/types';
import { useTimelineEventPreviewMethodsContext } from '@organization/src/components/Timeline/preview/context/TimelineEventPreviewContext';
import { Avatar } from '@ui/media/Avatar';
import { Flex } from '@ui/layout/Flex';
import { getName } from '@spaces/utils/getParticipantsName';
import { Button } from '@ui/form/Button';
import { DateTimeUtils } from '@spaces/utils/date';
import { InteractionEventWithDate } from '@organization/src/components/Timeline/types';
import { IntercomMessageCard } from '@organization/src/components/Timeline/events/intercom/IntercomMessageCard';
import { User02 } from '@ui/media/icons/User02';

// TODO unify with slack
export const IntercomStub: FC<{ intercomEvent: InteractionEventWithDate }> = ({
  intercomEvent,
}) => {
  const { openModal } = useTimelineEventPreviewMethodsContext();

  const intercomSender =
    (intercomEvent?.sentBy?.[0] as ContactParticipant)?.contactParticipant ||
    (intercomEvent?.sentBy?.[0] as JobRoleParticipant)?.jobRoleParticipant
      ?.contact ||
    (intercomEvent?.sentBy?.[0] as UserParticipant)?.userParticipant;
  const isSentByTenantUser =
    intercomEvent?.sentBy?.[0]?.__typename === 'UserParticipant';
  if (!intercomSender) {
    return null;
  }

  const intercomEventReplies = intercomEvent.interactionSession?.events?.filter(
    (e) => e?.id !== intercomEvent?.id,
  );
  const uniqThreadParticipants = intercomEventReplies
    ?.map((e) => {
      const threadSender =
        (e?.sentBy?.[0] as ContactParticipant)?.contactParticipant ||
        (e?.sentBy?.[0] as JobRoleParticipant)?.jobRoleParticipant?.contact ||
        (e?.sentBy?.[0] as UserParticipant)?.userParticipant;

      return threadSender;
    })
    ?.filter((v, i, a) => a.findIndex((t) => !!t && t?.id === v?.id) === i);

  return (
    <IntercomMessageCard
      name={getName(intercomSender)}
      profilePhotoUrl={intercomSender?.profilePhotoUrl}
      sourceUrl={intercomEvent?.externalLinks?.[0]?.externalUrl}
      content={intercomEvent?.content || ''}
      onClick={() => openModal(intercomEvent.id)}
      date={DateTimeUtils.formatTime(intercomEvent?.date)}
      showDateOnHover
      ml={isSentByTenantUser ? 6 : 0}
    >
      {!!intercomEventReplies?.length && (
        <Flex mt={1}>
          <Flex columnGap={1} mr={1}>
            {uniqThreadParticipants?.map(
              ({ id, name, firstName, lastName, profilePhotoUrl }) => {
                const displayName =
                  name ?? [firstName, lastName].filter(Boolean).join(' ');

                return (
                  <Avatar
                    size='xs'
                    name={displayName}
                    variant='roundedSquareSmall'
                    icon={<User02 color='primary.700' />}
                    src={profilePhotoUrl ? profilePhotoUrl : undefined}
                    key={`uniq-intercom-thread-participant-${intercomEvent.id}-${id}`}
                  />
                );
              },
            )}
          </Flex>
          <Button variant='link' fontSize='sm' size='sm'>
            {intercomEventReplies.length}{' '}
            {intercomEventReplies.length === 1 ? 'reply' : 'replies'}
          </Button>
        </Flex>
      )}
    </IntercomMessageCard>
  );
};
