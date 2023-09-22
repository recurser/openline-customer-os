import React, { useRef } from 'react';
import { CardHeader, CardBody } from '@ui/presentation/Card';
import { Heading } from '@ui/typography/Heading';
import { Text } from '@ui/typography/Text';
import { Flex } from '@ui/layout/Flex';
import { Tooltip } from '@ui/presentation/Tooltip';
import { IconButton } from '@ui/form/IconButton';
import { useTimelineEventPreviewContext } from '../../preview/TimelineEventsPreviewContext/TimelineEventPreviewContext';
import CopyLink from '@spaces/atoms/icons/CopyLink';
import Times from '@spaces/atoms/icons/Times';
import copy from 'copy-to-clipboard';
import { VStack } from '@ui/layout/Stack';
import { LogEntryWithAliases } from '@organization/components/Timeline/types';
import { User } from '@graphql/types';
import { Box } from '@ui/layout/Box';
import noteImg from 'public/images/note-img-preview.png';
import { LogEntryDatePicker } from './preview/LogEntryDatePicker';
import { Image } from '@ui/media/Image';
import { LogEntryExternalLink } from './preview/LogEntryExternalLink';
import { getGraphQLClient } from '@shared/util/getGraphQLClient';
import {
  LogEntryUpdateFormDto,
  LogEntryUpdateFormDtoI,
} from './preview/LogEntryUpdateFormDto';
import { useForm } from 'react-inverted-form';
import { useSession } from 'next-auth/react';
import { useUpdateLogEntryMutation } from '@organization/graphql/updateLogEntry.generated';
import { useQueryClient } from '@tanstack/react-query';
import { PreviewTags } from './preview/PreviewTags';
import { PreviewEditor } from './preview/PreviewEditor';

const getAuthor = (user: User) => {
  if (!user?.firstName && !user.lastName) {
    return 'Unknown';
  }

  return `${user.firstName} ${user.lastName}`.trim();
};

export const LogEntryPreviewModal: React.FC = () => {
  const { closeModal, modalContent } = useTimelineEventPreviewContext();
  const { data: session } = useSession();
  const event = modalContent as LogEntryWithAliases;
  const author = getAuthor(event?.logEntryCreatedBy);
  const authorEmail = event?.logEntryCreatedBy?.emails?.[0]?.email;
  const timeoutRef = useRef<NodeJS.Timeout | null>(null);

  const client = getGraphQLClient();
  const queryClient = useQueryClient();

  const formId = 'log-entry-update';
  const isAuthor =
    event.logEntryCreatedBy?.emails?.findIndex(
      (e) => session?.user?.email === e.email,
    ) !== -1;

  const updateLogEntryMutation = useUpdateLogEntryMutation(client, {
    onSuccess: () => {
      timeoutRef.current = setTimeout(
        () => queryClient.invalidateQueries(['GetTimeline.infinite']),
        500,
      );
    },
  });
  const logEntryStartedAtValues = new LogEntryUpdateFormDto(event);

  useForm<LogEntryUpdateFormDtoI>({
    formId,
    defaultValues: logEntryStartedAtValues,

    stateReducer: (state, action, next) => {
      if (action.type === 'FIELD_BLUR') {
        updateLogEntryMutation.mutate({
          id: event.id,
          input: {
            ...LogEntryUpdateFormDto.toPayload({
              ...state.values,
              [action.payload.name]: action.payload.value,
            }),
          },
        });
      }
      return next;
    },
  });

  return (
    <>
      <CardHeader
        py='4'
        px='6'
        pb='1'
        position='sticky'
        top={0}
        borderRadius='xl'
      >
        <Flex
          direction='row'
          justifyContent='space-between'
          alignItems='center'
        >
          <Flex alignItems='center'>
            <Heading size='sm' fontSize='lg'>
              Log entry
            </Heading>
          </Flex>
          <Flex direction='row' justifyContent='flex-end' alignItems='center'>
            <Tooltip label='Copy link' placement='bottom'>
              <IconButton
                variant='ghost'
                aria-label='Copy link to this entry'
                color='gray.500'
                fontSize='sm'
                size='sm'
                mr={1}
                icon={<CopyLink color='gray.500' height='18px' />}
                onClick={() => copy(window.location.href)}
              />
            </Tooltip>
            <Tooltip label='Close' aria-label='close' placement='bottom'>
              <IconButton
                variant='ghost'
                aria-label='Close preview'
                color='gray.500'
                fontSize='sm'
                size='sm'
                icon={<Times color='gray.500' height='24px' />}
                onClick={closeModal}
              />
            </Tooltip>
          </Flex>
        </Flex>
      </CardHeader>
      <CardBody
        mt={0}
        maxHeight='calc(100vh - 9rem)'
        p={6}
        pt={0}
        overflow='auto'
      >
        <Box position='relative'>
          <Image
            src={noteImg}
            alt=''
            height={123}
            width={174}
            position='absolute'
            top={-2}
            right={-3}
          />
        </Box>
        <VStack gap={2} alignItems='flex-start'>
          <Flex direction='column'>
            <LogEntryDatePicker event={event} formId={formId} />
          </Flex>
          <Flex direction='column'>
            <Text fontSize='sm' fontWeight='semibold'>
              Author
            </Text>
            <Tooltip label={authorEmail} hasArrow>
              <Text fontSize='sm'>{author}</Text>
            </Tooltip>
          </Flex>

          <Flex direction='column' w='full'>
            <Text fontSize='sm' fontWeight='semibold'>
              Entry
            </Text>

            <PreviewEditor
              isAuthor={isAuthor}
              formId={formId}
              initialContent={`${event?.content}`}
              tags={[]}
            />
          </Flex>

          <PreviewTags isAuthor={isAuthor} tags={event.tags} formId={formId} />

          {event?.externalLinks?.[0]?.externalUrl && (
            <LogEntryExternalLink externalLink={event?.externalLinks?.[0]} />
          )}
        </VStack>
      </CardBody>
    </>
  );
};