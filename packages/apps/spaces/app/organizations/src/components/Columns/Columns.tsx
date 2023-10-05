import { Flex } from '@ui/layout/Flex';
import { Organization } from '@graphql/types';
import { THead, createColumnHelper } from '@ui/presentation/Table';
import { Skeleton, SkeletonCircle } from '@ui/presentation/Skeleton';

import { OwnerCell } from './Cells/owner/OwnerCell';
import { RelationshipStage } from './Cells/stage/RelationshipStage';
import { TimeToRenewalCell } from './Cells/renewal/TimeToRenewalCell';
import { OrganizationCell } from './Cells/organization/OrganizationCell';
import { RenewalForecastCell } from './Cells/renewal/RenewalForecastCell';
import { LastTouchpointCell } from './Cells/touchpoint/LastTouchpointCell';
import { RenewalLikelihoodCell } from './Cells/renewal/RenewalLikelihoodCell';
import { OrganizationRelationship } from './Cells/relationship/OrganizationRelationship';

const columnHelper =
  createColumnHelper<Omit<Organization, 'lastTouchPointTimelineEvent'>>();

export const columns = (tabs?: { [key: string]: string } | null) => [
  columnHelper.accessor((row) => row, {
    id: 'ORGANIZATION',
    cell: (props) => {
      return (
        <OrganizationCell
          key={props.getValue().id}
          organization={props.getValue()}
          lastPositionParams={tabs?.[props.getValue()?.id]}
        />
      );
    },
    minSize: 200,
    header: (props) => <THead<Organization> title='Company' {...props} />,
    skeleton: () => (
      <Flex align='center' h='full'>
        <SkeletonCircle size='48px' startColor='gray.300' endColor='gray.300' />
        <Flex ml='2' flexDir='column' h='42px' align='center' gap='1'>
          <Skeleton
            width='100px'
            height='18px'
            startColor='gray.300'
            endColor='gray.300'
          />
          <Skeleton
            width='100px'
            height='18px'
            startColor='gray.300'
            endColor='gray.300'
          />
        </Flex>
      </Flex>
    ),
  }),
  columnHelper.accessor('relationshipStages', {
    id: 'RELATIONSHIP',
    header: (props) => (
      <THead<Organization> title='Relationship | Stage' {...props} />
    ),
    minSize: 200,
    cell: (props) => {
      const relationshipStages = props.getValue();
      const relationship = relationshipStages?.[0]?.relationship;
      const stage = relationshipStages?.[0]?.stage;
      const organizationId = props.row.original.id;

      return (
        <>
          <OrganizationRelationship
            defaultValue={relationship}
            organizationId={organizationId}
          />
          <RelationshipStage
            defaultValue={stage}
            relationship={relationship}
            organizationId={organizationId}
          />
        </>
      );
    },
    skeleton: () => (
      <Flex gap='1' flexDir='column'>
        <Skeleton
          width='100%'
          height='18px'
          startColor='gray.300'
          endColor='gray.300'
        />
        <Skeleton
          width='25%'
          height='18px'
          startColor='gray.300'
          endColor='gray.300'
        />
      </Flex>
    ),
  }),
  columnHelper.accessor('accountDetails', {
    id: 'RENEWAL_LIKELIHOOD',
    minSize: 200,
    cell: (props) => {
      const organizationId = props.row.original.id;
      const value = props.getValue()?.renewalLikelihood;
      const currentProbability = value?.probability;
      const previousProbability = value?.previousProbability;
      const updatedAt = value?.updatedAt;

      return (
        <RenewalLikelihoodCell
          updatedAt={updatedAt}
          organizationId={organizationId}
          currentProbability={currentProbability}
          previousProbability={previousProbability}
        />
      );
    },
    header: (props) => (
      <THead<Organization> title='Renewal Likelihood' {...props} />
    ),
    skeleton: () => (
      <Flex flexDir='column' gap='1'>
        <Skeleton
          width='25%'
          height='18px'
          startColor='gray.300'
          endColor='gray.300'
        />
        <Skeleton
          width='75%'
          height='18px'
          startColor='gray.300'
          endColor='gray.300'
        />
      </Flex>
    ),
  }),
  columnHelper.accessor('accountDetails', {
    id: 'RENEWAL_CYCLE_NEXT',
    minSize: 200,
    cell: (props) => {
      const values = props.getValue()?.billingDetails;
      const renewalDate = values?.renewalCycleNext;
      const renewalFrequency = values?.renewalCycle;

      return (
        <TimeToRenewalCell
          renewalDate={renewalDate}
          renewalFrequency={renewalFrequency}
        />
      );
    },
    header: (props) => (
      <THead<Organization> title='Time to Renewal' {...props} />
    ),
    skeleton: () => (
      <Skeleton
        width='50%'
        height='18px'
        startColor='gray.300'
        endColor='gray.300'
      />
    ),
  }),
  columnHelper.accessor('accountDetails', {
    id: 'FORECAST_AMOUNT',
    minSize: 200,
    cell: (props) => {
      const value = props.getValue()?.renewalForecast;
      const amount = value?.amount;
      const potentialAmount = value?.potentialAmount;

      return (
        <RenewalForecastCell
          amount={amount}
          potentialAmount={potentialAmount}
          isUpdatedByUser={!!value?.updatedById}
        />
      );
    },
    header: (props) => (
      <THead<Organization> title='Renewal Forecast' {...props} />
    ),
    skeleton: () => (
      <Flex flexDir='column' gap='1'>
        <Skeleton
          width='50%'
          height='18px'
          startColor='gray.300'
          endColor='gray.300'
        />
        <Skeleton
          width='25%'
          height='18px'
          startColor='gray.300'
          endColor='gray.300'
        />
      </Flex>
    ),
  }),
  columnHelper.accessor('owner', {
    id: 'OWNER',
    minSize: 200,
    cell: (props) => (
      <OwnerCell id={props.row.original.id} owner={props.getValue()} />
    ),
    header: (props) => <THead<Organization> title='Owner' {...props} />,
    skeleton: () => (
      <Skeleton
        width='75%'
        height='18px'
        startColor='gray.300'
        endColor='gray.300'
      />
    ),
  }),
  columnHelper.accessor('market', {
    id: 'LAST_TOUCHPOINT',
    minSize: 200,
    cell: (props) => (
      <LastTouchpointCell
        lastTouchPointAt={props.row.original.lastTouchPointAt}
        lastTouchPointTimelineEvent={
          (props.row.original as Organization).lastTouchPointTimelineEvent
        }
      />
    ),
    header: (props) => (
      <THead<Organization> title='Last Touchpoint' {...props} />
    ),
    skeleton: () => (
      <Flex flexDir='column' gap='1'>
        <Skeleton
          width='75%'
          height='18px'
          startColor='gray.300'
          endColor='gray.300'
        />
        <Skeleton
          width='100%'
          height='18px'
          startColor='gray.300'
          endColor='gray.300'
        />
      </Flex>
    ),
  }),
];