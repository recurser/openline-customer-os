import React from 'react';
import { Box } from '@ui/layout/Box';
import { LogEntryTimelineAction } from './logger/LogEntryTimelineAction';
import { useTimelineActionContext } from '@organization/src/components/Timeline/TimelineActions/context/TimelineActionContext';
import { EmailTimelineAction } from './email/EmailTimelineAction';

export const TimelineActionsArea: React.FC = () => {
  const { openedEditor } = useTimelineActionContext();

  return (
    <Box
      bg={'#F9F9FB'}
      borderTop='1px dashed'
      borderTopColor='gray.200'
      pt={openedEditor !== null ? 6 : 0}
      pb={openedEditor !== null ? 2 : 8}
      mt={-4}
    >
      <EmailTimelineAction />
      <LogEntryTimelineAction />
    </Box>
  );
};
