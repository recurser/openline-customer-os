import React, { useEffect, useRef } from 'react';
import { Text } from '@ui/typography/Text';
import { TagsSelect } from '@organization/src/components/Timeline/TimelineActions/logger/components/TagSelect';
import { useForm } from 'react-inverted-form';
import {
  LogEntryTagsDto,
  LogEntryTagsFormDtoI,
} from '@organization/src/components/Timeline/events/logEntry/preview/tags/LogEntryTagsDto';
import { Tag } from '@graphql/types';
import { useQueryClient } from '@tanstack/react-query';
import { getGraphQLClient } from '@shared/util/getGraphQLClient';
import { useResetLogEntryTagsMutation } from '@organization/src/graphql/resetLogEntryTags.generated';

export const PreviewTags: React.FC<{
  isAuthor: boolean;
  tags?: Array<Tag>;
  id: string;
  tagOptions?: Array<{ label: string; value: string }>;
}> = ({ isAuthor, tags = [], id, tagOptions }) => {
  const logEntryStartedAtValues = new LogEntryTagsDto({ tags });
  const formId = 'preview-modal-log-entry-tag-update';
  const queryClient = useQueryClient();
  const client = getGraphQLClient();
  const timeoutRef = useRef<NodeJS.Timeout | null>(null);
  const updateLogEntryTags = useResetLogEntryTagsMutation(client, {
    onSuccess: () => {
      timeoutRef.current = setTimeout(
        () => queryClient.invalidateQueries(['GetTimeline.infinite']),
        500,
      );
    },
  });
  useForm<LogEntryTagsFormDtoI>({
    formId,
    defaultValues: logEntryStartedAtValues,

    stateReducer: (state, action, next) => {
      if (action.type === 'FIELD_CHANGE') {
        updateLogEntryTags.mutate({
          id: id,
          input: [
            ...LogEntryTagsDto.toPayload({
              tags: action.payload.value,
            }).tags,
          ],
        });
      }
      return next;
    },
  });

  useEffect(() => {
    return () => {
      if (timeoutRef.current) {
        clearTimeout(timeoutRef.current);
      }
    };
  }, []);

  return (
    <>
      {!isAuthor && (
        <Text fontSize='sm' fontWeight='medium'>
          {tags.map(({ name }) => `#${name}`).join(' ')}
        </Text>
      )}

      {isAuthor && (
        <Text
          fontSize='sm'
          fontWeight='medium'
          lineHeight='1'
          sx={{
            '--tag-select-font-size': `14px`,
          }}
        >
          <TagsSelect formId={formId} name='tags' tags={tagOptions} />
        </Text>
      )}
    </>
  );
};
