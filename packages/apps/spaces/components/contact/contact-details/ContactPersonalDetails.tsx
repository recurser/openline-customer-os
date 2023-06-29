import React, { FC, useEffect, useState } from 'react';
import { EditableContentInput } from '@spaces/atoms/input';
import { DeleteConfirmationDialog } from '@spaces/atoms/delete-confirmation-dialog';
import styles from './contact-details.module.scss';
import {
  useArchiveContact,
  useUpdateContactPersonalDetails,
} from '@spaces/hooks/useContact';
import { ContactDetailsSkeleton } from './skeletons';
import { ContactTags } from '../contact-tags';
import { ContactAvatar } from '../../ui-kit/molecules/organization-avatar';
import { useRecoilValue } from 'recoil';
import { contactDetailsEdit } from '../../../state';
import { JobRoleInput } from './edit';
import { IconButton } from '@spaces/atoms/icon-button/IconButton';
import Inbox from '@spaces/atoms/icons/Inbox';
import classNames from 'classnames';
import { useCreateContactJobRole } from '@spaces/hooks/useContactJobRole/useCreateContactJobRole';
import { useAutoAnimate } from '@formkit/auto-animate/react';
import { getContactDisplayName } from '../../../utils';
import { Contact } from '@spaces/graphql';
import { ContactDetailsProps } from '@spaces/contact/contact-details/type';

export const ContactPersonalDetails: FC<ContactDetailsProps> = ({
  id,
  data,
  loading,
}) => {
  const { isEditMode } = useRecoilValue(contactDetailsEdit);
  const { onCreateContactJobRole } = useCreateContactJobRole({ contactId: id });

  const { onUpdateContactPersonalDetails } = useUpdateContactPersonalDetails({
    contactId: id,
  });
  const [deleteConfirmationModalVisible, setDeleteConfirmationModalVisible] =
    useState(false);
  const { onArchiveContact } = useArchiveContact({ id });
  const [animatedJobRolesRowParent] = useAutoAnimate({
    easing: 'linear',
  });

  useEffect(() => {
    if (!loading && !data?.jobRoles?.length && isEditMode) {
      onCreateContactJobRole({ jobTitle: '', primary: true });
    }
  }, [loading, data?.jobRoles.length, isEditMode]);

  if (loading) {
    return <ContactDetailsSkeleton />;
  }

  return (
    <div className={styles.header}>
      <div className={styles.avatarWrapper}>
        <div className={styles.photo}>
          <ContactAvatar
            size={50}
            name={getContactDisplayName(data as Contact)}
          />
        </div>
        {isEditMode && (
          <>
            <IconButton
              label='Archive'
              className={styles.archiveContactButton}
              size='xxxs'
              mode='danger'
              onClick={() => setDeleteConfirmationModalVisible(true)}
              icon={<Inbox height={20} />}
            />
            <DeleteConfirmationDialog
              deleteConfirmationModalVisible={deleteConfirmationModalVisible}
              setDeleteConfirmationModalVisible={
                setDeleteConfirmationModalVisible
              }
              deleteAction={() =>
                onArchiveContact().then(() =>
                  setDeleteConfirmationModalVisible(false),
                )
              }
              header='Confirm archive'
              confirmationButtonLabel='Archive contact'
              explanationText='Are you sure you want to archive this contact?'
            />
          </>
        )}
      </div>
      <div className={styles.name}>
        <div className={styles.nameAndEditButton}>
          <div className={styles.nameContainer}>
            <EditableContentInput
              id={`conatct-personal-details-first-name-${id}`}
              label='First name'
              isEditMode={isEditMode}
              value={data?.firstName || data?.name || ''}
              placeholder={isEditMode ? 'First name' : 'Unnamed'}
              onBlur={(value: string) =>
                onUpdateContactPersonalDetails({
                  firstName: value,
                  lastName: data?.lastName || '',
                })
              }
            />
            <EditableContentInput
              id={`conatct-personal-details-last-name-${id}`}
              label='Last name'
              isEditMode={isEditMode}
              value={data?.lastName || ''}
              placeholder={isEditMode ? 'Last name' : ''}
              onBlur={(value: string) => {
                return onUpdateContactPersonalDetails({
                  lastName: value,
                  firstName: data?.firstName || '',
                });
              }}
            />
          </div>
        </div>

        <div ref={animatedJobRolesRowParent}>
          {data?.jobRoles?.map((jobRole: any, index: number) => {
            return (
              <JobRoleInput
                key={jobRole.id}
                contactId={id}
                organization={jobRole.organization}
                primary={jobRole.primary}
                jobRole={jobRole?.jobTitle || ''}
                roleId={jobRole.id}
                isEditMode={isEditMode}
                showAddButton={
                  data?.jobRoles.length
                    ? data.jobRoles.length - 1 === index
                    : true
                }
              />
            );
          })}
        </div>

        <ContactTags
          id={id}
          mode={isEditMode ? 'EDIT' : 'PREVIEW'}
          tags={data?.tags}
        />
        <div
          className={classNames(styles.source, {
            [styles.sourceEditMode]: isEditMode,
          })}
        >
          Source:
          <div>{data?.source || 'OPENLINE'}</div>
        </div>
      </div>
    </div>
  );
};
