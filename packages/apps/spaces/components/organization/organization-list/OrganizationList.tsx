'use client';

import React, {
  useEffect,
  useMemo,
  useState,
  lazy,
  Suspense,
  useCallback,
} from 'react';
import styles from './organization-list.module.scss';
import { columns } from './OrganizationListColumns';
import { useFinderOrganizationTableData } from '@spaces/hooks/useFinderOrganizationTableData';
import { useGCliSearch } from '@spaces/hooks/useGCliSearch';
import { GCLIContextProvider, GCLIInput } from '@spaces/molecules/gCLI';

import {
  Table,
  RowSelectionState,
  TableInstance,
  SortingState,
} from '@ui/presentation/Table';

import {
  useCreateOrganization,
  useMergeOrganizations,
} from '@spaces/hooks/useOrganization';
import {
  SortingDirection,
  type Filter,
  type Organization,
  type SortBy,
} from '@graphql/types';

import { mapGCliSearchTermsToFilterList } from '@spaces/utils/mapGCliSearchTerms';
import { useRouter } from 'next/router';

import { useDisclosure } from '@chakra-ui/react-use-disclosure';
import { ConfirmDeleteDialog } from '@ui/overlay/AlertDialog/ConfirmDeleteDialog';
import { Text } from '@chakra-ui/react';
import { useLocalStorage, useReadLocalStorage } from 'usehooks-ts';

const OrganizationListActions = lazy(() => import('./OrganizationListActions'));
import EmptyState from './EmptyState';
import { Archive } from '@ui/media/icons/Archive';
import { useHideOrganizations } from '@spaces/hooks/useOrganization/useHideOrganizations';

interface OrganizationListProps {
  preFilters?: Array<Filter>;
  label: string;
  filterLabel: string;
  icon: React.ReactNode;
}

export const OrganizationList: React.FC<OrganizationListProps> = ({
  preFilters,
  filterLabel,
  label,
  icon,
}: OrganizationListProps) => {
  const { isOpen, onOpen, onClose } = useDisclosure();
  const tabs = useReadLocalStorage<{ [key: string]: string }>(
    `customeros-player-last-position`,
  );
  const [idsToRemove, setIdsToRemove] = useState<Array<string>>([]);

  const [tableInstance, setTableInstance] =
    useState<TableInstance<Organization> | null>(null);
  const [page, setPagination] = useState(1);
  const [sorting, setSorting] = useState<SortingState>([]);
  const [enableSelection, setEnableSelection] = useState(false);
  const [selection, setSelection] = useState<RowSelectionState>({});
  const [suggestions, setSuggestions] = useState<any[]>([]);
  const { data: gcliData, loading: gcliLoading, refetch } = useGCliSearch();
  const sortBy: SortBy | undefined = useMemo(() => {
    setPagination(1);
    if (!sorting.length) return;
    return {
      by: sorting[0].id,
      direction: sorting[0].desc ? SortingDirection.Desc : SortingDirection.Asc,
      caseSensitive: false,
    };
  }, [sorting]);

  const { push } = useRouter();

  const { onMergeOrganizations } = useMergeOrganizations();
  const { onHideOrganizations } = useHideOrganizations();
  const { onCreateOrganization } = useCreateOrganization();

  const [organizationFilters, setOrganizationFilters] = useLocalStorage<
    Filter[]
  >(`customeros-${filterLabel}-list-filters`, []);
  const { data, loading, fetchMore, variables, totalElements } =
    useFinderOrganizationTableData(
      preFilters
        ? [...organizationFilters, ...preFilters]
        : organizationFilters,
      sortBy,
    );
  const handleFilterResults = (searchTerms: any[]) => {
    setPagination(1);

    let filters = mapGCliSearchTermsToFilterList(searchTerms, 'ORGANIZATION');
    setOrganizationFilters(filters);
    if (preFilters) {
      filters = [...filters, ...preFilters];
    }
    fetchMore({
      variables: {
        pagination: {
          page: 1,
          limit: 20,
        },
        where: {
          AND: filters,
        },
        sort: sortBy,
      },
    });
  };

  const handleCreateOrganization = async () => {
    const newOrganization = await onCreateOrganization({ name: '' });
    if (newOrganization?.id) {
      push(`/organizations/${newOrganization?.id}`);
    }
  };

  const handleFetchMore = useCallback(() => {
    setPagination((prev) => {
      return prev + 1;
    });
    fetchMore({
      variables: {
        pagination: {
          limit: variables.pagination.limit,
          page: page + 1,
        },
      },
    });
  }, [page, fetchMore, variables.pagination.limit]);

  const handleMergeOrganizations = (table: TableInstance<Organization>) => {
    const organizationIds = Object.keys(selection)
      .map((key) => data?.[Number(key)]?.id)
      .filter(Boolean) as string[];

    const primaryId = organizationIds[0];
    const mergeIds = organizationIds.slice(1);

    onMergeOrganizations({
      primaryOrganizationId: primaryId,
      mergedOrganizationIds: mergeIds,
      onSuccess: () => {
        setEnableSelection(false);
        table.resetRowSelection();
      },
    });
  };
  const handleHideOrganizations = () => {
    onHideOrganizations({
      ids: idsToRemove,
    });
    onClose();
    tableInstance?.resetRowSelection();
    setEnableSelection(false);
    setIdsToRemove([]);
    setTableInstance(null);
  };
  const handleCancelRemoveOrganizations = () => {
    onClose();
    tableInstance?.resetRowSelection();
    setEnableSelection(false);
    setTableInstance(null);
  };
  const handleOpenConfirmationModal = (table: TableInstance<Organization>) => {
    const organizationIds = Object.keys(selection)
      .map((key) => data?.[Number(key)]?.id)
      .filter(Boolean) as string[];
    setIdsToRemove(organizationIds);
    setTableInstance(table);
    onOpen();
  };

  useEffect(() => {
    if (!gcliLoading && gcliData) {
      setSuggestions(gcliData);
    }
  }, [gcliLoading, gcliData]);

  const existingTerms = useMemo(() => {
    return organizationFilters?.reduce(
      (
        acc: { type: string; display: string; highlighted: boolean }[],
        item,
      ) => {
        if (item?.filter?.property === 'ORGANIZATION') {
          acc.push({
            type: 'GENERIC',
            display: item?.filter?.value,
            highlighted: false,
          });
        }
        return acc;
      },
      [],
    );
  }, [organizationFilters]);

  if (totalElements === 0) {
    return <EmptyState onClick={handleCreateOrganization} />;
  }

  return (
    <>
      <div className={styles.inputSection}>
        <GCLIContextProvider
          label={label}
          icon={icon}
          existingTerms={existingTerms}
          loadSuggestions={(searchTerm: string) => {
            refetch && refetch({ limit: 5, keyword: searchTerm });
          }}
          loadingSuggestions={gcliLoading}
          suggestionsLoaded={suggestions}
          onItemsChange={handleFilterResults}
          selectedTermFormat={(item: any) => {
            if (item.type === 'STATE') {
              return item.data[0].value;
            }
            return item.display;
          }}
        >
          <GCLIInput />
        </GCLIContextProvider>
        {totalElements && (
          <Text
            color='gray.500'
            fontSize='xs'
            whiteSpace='nowrap'
            ml={3}
            alignSelf='center'
          >
            Total items: {totalElements}
          </Text>
        )}
      </div>

      <ConfirmDeleteDialog
        label={`Archive selected ${
          idsToRemove.length === 1 ? 'organization' : 'organizations'
        }?`}
        confirmButtonLabel={'Archive'}
        isOpen={isOpen}
        onClose={handleCancelRemoveOrganizations}
        onConfirm={handleHideOrganizations}
        isLoading={loading}
        icon={<Archive />}
      />

      {/* TODO: Remove coercion to any type when we get rid of the old graphql types generated which are out of sync */}
      <Table<Organization>
        data={(data as unknown as Organization[]) ?? []}
        columns={columns(tabs)}
        sorting={sorting}
        enableTableActions
        isLoading={loading}
        selection={selection}
        onSortingChange={setSorting}
        onFetchMore={handleFetchMore}
        totalItems={totalElements ?? 0}
        onSelectionChange={setSelection}
        enableRowSelection={enableSelection}
        renderTableActions={(table) => (
          <Suspense fallback={<div />}>
            <OrganizationListActions
              table={table as any}
              selection={selection}
              isSelectionEnabled={enableSelection}
              toggleSelection={setEnableSelection}
              onCreateOrganization={handleCreateOrganization}
              onMergeOrganizations={handleMergeOrganizations as any}
              onArchiveOrganizations={handleOpenConfirmationModal as any}
            />
          </Suspense>
        )}
      />
    </>
  );
};
