import React from 'react';
import styles from './health-indicator-owner.module.scss';
import { HealthIndicatorSelect } from '@spaces/organization/health-select/HealthIndicatorSelect';
import { HealthIndicator } from '@spaces/graphql';

interface OrganizationOwnerProps {
  id: string;
  healthIndicator: HealthIndicator | undefined | null;
}

export const OrganizationHealthIndicator: React.FC<OrganizationOwnerProps> = ({
  id,
  healthIndicator,
}) => {
  return (
    <article className={styles.health_section}>
      <h1 className={styles.health_header}>Health</h1>

      <HealthIndicatorSelect
        organizationId={id}
        showIcon
        healthIndicator={healthIndicator}
      />
    </article>
  );
};