import React, { useEffect, useRef } from 'react';
import { Box } from '@ui/layout/Box';
import { ComposeEmailContainer } from '@organization/src/components/Timeline/events/email/compose-email/ComposeEmailContainer';
import { useTimelineActionEmailContext } from '@organization/src/components/Timeline/TimelineActions/context/TimelineActionEmailContext';
import { useTimelineActionContext } from '@organization/src/components/Timeline/TimelineActions/context/TimelineActionContext';
import { KeymapperClose } from '@ui/form/RichTextEditor/components/keyboardShortcuts/KeymapperClose';
import { useTimelineRefContext } from '@organization/src/components/Timeline/context/TimelineRefContext';

export const EmailTimelineAction: React.FC = () => {
  const {
    remirrorProps,
    isSending,
    onCreateEmail,
    formId,
    state,
    checkCanExitSafely,
  } = useTimelineActionEmailContext();
  const { virtuosoRef } = useTimelineRefContext();
  const { openedEditor, showEditor } = useTimelineActionContext();

  const isEmail = openedEditor === 'email';
  const emailWrapperRef = useRef(null);

  useEffect(() => {
    if (isEmail) {
      virtuosoRef?.current?.scrollBy({ top: 300 });
    }
  }, [isEmail, virtuosoRef]);

  const handleClose = () => {
    const canClose = checkCanExitSafely();

    if (canClose) {
      showEditor(null);
    }
  };

  return (
    <>
      {isEmail && (
        <Box
          ref={emailWrapperRef}
          borderRadius={'md'}
          boxShadow={'lg'}
          m={6}
          mt={2}
          bg={'white'}
          border='1px solid'
          borderColor='gray.100'
        >
          <ComposeEmailContainer
            formId={formId}
            modal={false}
            onClose={handleClose}
            to={state.values.to}
            cc={state.values.cc}
            bcc={state.values.bcc}
            onSubmit={onCreateEmail}
            isSending={isSending}
            remirrorProps={remirrorProps}
          >
            <KeymapperClose onClose={handleClose} />
          </ComposeEmailContainer>
        </Box>
      )}
    </>
  );
};
