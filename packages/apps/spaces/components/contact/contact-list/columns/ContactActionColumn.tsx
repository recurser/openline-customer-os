import React, { useEffect } from 'react';
import { useMergeContacts } from '@spaces/hooks/useContact';
import { ActionColumn } from '@spaces/finder/finder-table';
import { useSetRecoilState } from 'recoil';
import { tableMode } from '@spaces/finder/state';
import { useRouter } from 'next/router';

export const ContactActionColumn: React.FC = () => {
  const { onMergeContacts } = useMergeContacts();
  const setTableMode = useSetRecoilState(tableMode);
  const { push } = useRouter();

  useEffect(() => {
    return () => {
      setTableMode('PREVIEW');
    };
  }, []);

  return (
    <ActionColumn
      onMerge={({ primaryId, mergeIds }) =>
        onMergeContacts({
          primaryContactId: primaryId,
          mergedContactIds: mergeIds,
        })
      }
      actions={[
        {
          label: 'Add Contact',
          command() {
            push('/contact/new');
          },
        },
        {
          label: 'Merge Contacts',
          command() {
            return setTableMode('MERGE');
          },
        },
      ]}
    />
  );
};
