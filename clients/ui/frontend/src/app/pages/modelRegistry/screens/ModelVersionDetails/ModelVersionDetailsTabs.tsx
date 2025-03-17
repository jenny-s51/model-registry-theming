import * as React from 'react';
import { useNavigate } from 'react-router-dom';
import { Button, Content, PageSection, Tab, Tabs, TabTitleText } from '@patternfly/react-core';
import { ModelVersion } from '~/app/types';
import { ModelVersionDetailsTabTitle, ModelVersionDetailsTab } from './const';
import ModelVersionDetailsView from './ModelVersionDetailsView';
// import EmptyModelRegistryState from '../components/EmptyModelRegistryState';
// TODO: [Model Serving] Uncomment when model serving is available
// import ModelVersionRegisteredDeploymentsView from './ModelVersionsRegisteredDeploymentsView';

type ModelVersionDetailTabsProps = {
  tab: ModelVersionDetailsTab;
  modelVersion: ModelVersion;
  isArchiveVersion?: boolean;
  refresh: () => void;
};

const ModelVersionDetailsTabs: React.FC<ModelVersionDetailTabsProps> = ({
  tab,
  modelVersion: mv,
  isArchiveVersion,
  refresh,
}) => {
  const navigate = useNavigate();
  return (
    <Tabs
      activeKey={tab}
      aria-label="Model versions details page tabs"
      role="region"
      data-testid="model-versions-details-page-tabs"
      onSelect={(_event, eventKey) => navigate(`../${eventKey}`, { relative: 'path' })}
    >
      <Tab
        eventKey={ModelVersionDetailsTab.DETAILS}
        title={<TabTitleText>{ModelVersionDetailsTabTitle.DETAILS}</TabTitleText>}
        aria-label="Model versions details tab"
        data-testid="model-versions-details-tab"
      >
        <PageSection
          hasBodyWrapper={false}
          isFilled
          data-testid="model-versions-details-tab-content"
        >
          <ModelVersionDetailsView
            modelVersion={mv}
            refresh={refresh}
            isArchiveVersion={isArchiveVersion}
          />
        </PageSection>
      </Tab>
      <Tab
        eventKey={ModelVersionDetailsTab.DEPLOYMENTS}
        title={<TabTitleText>{ModelVersionDetailsTabTitle.DEPLOYMENTS}</TabTitleText>}
        aria-label="Model versions deployments tab"
        data-testid="model-versions-deployments-tab"
      >
        <PageSection hasBodyWrapper isFilled data-testid="model-versions-deployments-tab-content">
          <ModelVersionDetailsView
            modelVersion={mv}
            refresh={refresh}
            isArchiveVersion={isArchiveVersion}
          />
        </PageSection>
      </Tab>
    </Tabs>
  );
};

export default ModelVersionDetailsTabs;
