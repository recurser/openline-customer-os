import React, { useRef, useState } from 'react';
import sanitizeHtml from 'sanitize-html';
import styles from './email-timeline-item.module.scss';
import { Button } from '../../atoms';
import linkifyHtml from 'linkify-html';
import { EmailParticipants } from './email-participants';
import classNames from 'classnames';
import { useContactCommunicationChannelsDetails } from '../../../../hooks/useContact';

interface Props {
  content: string;
  contentType: string;
  sentBy: Array<any>;
  sentTo: Array<any>;
  interactionSession: any;
  contactId?: string;
  isToDeprecate?: boolean; //remove
  deprecatedCC?: any; //remove
  deprecatedBCC?: any; //remove
}

export const EmailTimelineItem: React.FC<Props> = ({
  content,
  contentType,
  sentBy,
  sentTo,
  interactionSession,
  isToDeprecate = false,
  contactId,
  deprecatedCC,
  deprecatedBCC,
  ...rest
}) => {
  const { data, loading, error } = useContactCommunicationChannelsDetails({
    id: contactId || '',
  });
  const sentByExist =
    sentBy &&
    sentBy.length > 0 &&
    sentBy[0].__typename === 'EmailParticipant' &&
    sentBy[0].emailParticipant;
  const from = sentByExist ? sentBy[0].emailParticipant.email : '';
  const to =
    sentTo && sentTo.length > 0
      ? sentTo
          .filter((p: any) => p.type === 'TO')
          .map((p: any) => {
            if (p.__typename === 'EmailParticipant' && p.emailParticipant) {
              return p.emailParticipant.email;
            }
            return '';
          })
          .join('; ')
      : '';

  const cc =
    sentTo && sentTo.length > 0
      ? sentTo
          .filter((p: any) => p.type === 'CC')
          .map((p: any) => {
            if (p.__typename === 'EmailParticipant' && p.emailParticipant) {
              return p.emailParticipant.email;
            }
            return '';
          })
          .join('; ')
      : '';

  const bcc =
    sentTo && sentTo.length > 0
      ? sentTo
          .filter((p: any) => p.type === 'BCC')
          .map((p: any) => {
            if (p.__typename === 'EmailParticipant' && p.emailParticipant) {
              return p.emailParticipant.email;
            }
            return '';
          })
          .join('; ')
      : '';

  const [expanded, toggleExpanded] = useState(false);
  const timelineItemRef = useRef<HTMLDivElement>(null);

  const handleToggleExpanded = () => {
    toggleExpanded(!expanded);
    if (timelineItemRef?.current && expanded) {
      timelineItemRef?.current?.scrollIntoView({ behavior: 'smooth' });
    }
  };

  const isSentByContact =
    !!contactId &&
    !error &&
    !loading &&
    data?.emails.findIndex(({ email }) => email === from) !== -1;

  return (
    <div
      className={classNames({
        [styles.sendBy]: isSentByContact,
        [styles.sendTo]: !isSentByContact,
      })}
    >
      <div
        className={classNames(styles.emailWrapper, {
          [styles.expanded]: expanded,
        })}
      >
        <div ref={timelineItemRef} className={styles.scrollToView} />
        <article className={`${styles.emailContainer}`}>
          <div>
            <EmailParticipants
              from={isToDeprecate ? sentBy : from}
              to={isToDeprecate ? sentTo : new Array(to)}
              subject={interactionSession?.name}
              cc={isToDeprecate ? deprecatedCC : cc}
              bcc={isToDeprecate ? deprecatedBCC : bcc}
            />
          </div>

          <div
            className={`${styles.emailContentContainer} ${
              !expanded ? styles.eclipse : ''
            }`}
          >
            {contentType === 'text/html' && (
              <div
                className={`text-overflow-ellipsis ${styles.emailContent}`}
                dangerouslySetInnerHTML={{
                  __html: sanitizeHtml(
                    linkifyHtml(content, {
                      defaultProtocol: 'https',
                      rel: 'noopener noreferrer',
                    }),
                  ),
                }}
              ></div>
            )}
            {contentType === 'text/plain' && (
              <div className={`text-overflow-ellipsis ${styles.emailContent}`}>
                {content}
              </div>
            )}

            {!expanded && <div className={styles.eclipse} />}
          </div>
        </article>
      </div>
      <div className={styles.folderTab}>
        <Button
          onClick={() => handleToggleExpanded()}
          mode='link'
          className={styles.toggleExpandButton}
        >
          {expanded ? 'Collapse' : 'Expand'}
        </Button>
      </div>
    </div>
  );
};
