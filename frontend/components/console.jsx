import _ from 'lodash';
import {Row, Tabs, Tab} from 'react-bootstrap';
import React, {Component} from 'react';

import {alertError} from '../lib/error';
import {requestJson} from '../lib/request';
import LayoutView from './layout';
import ApplicationsTab from './console/application';
import DataSourcesTab from './console/data-source';

export default class ConsoleView extends Component {
  constructor(props) {
    super(props);

    this.state = {
      apps: [],
      dataSources: []
    };
  }

  componentDidMount() {
    return Promise.all([
      requestJson('/apps').then( apps => {
        this.setState({apps});
      }).catch(alertError),

      requestJson('/data-sources').then( dataSources => {
        this.setState({dataSources});
      }).catch(alertError)
    ]);
  }

  render() {
    const activeKey = _.trimStart(_.trimStart(this.props.location.pathname, this.props.match.path), '/');

    return <LayoutView>
      <Row>
        <Tabs animation={false} activeKey={activeKey ? activeKey : 'applications'} onSelect={::this.onTabSelected} id='console-tabs'>
          <Tab eventKey='applications' title='Applications'>
            <ApplicationsTab apps={this.state.apps} onAppEdited={::this.onAppEdited} onAppDeleted={::this.onAppDeleted} />
          </Tab>
          <Tab eventKey='data-sources' title='Data Sources'>
            <DataSourcesTab apps={this.state.apps} dataSources={this.state.dataSources} onDataSourceEdited={::this.onDataSourceEdited} onDataSourceDeleted={::this.onDataSourceDeleted} />
          </Tab>
        </Tabs>
      </Row>
    </LayoutView>;
  }

  onTabSelected(tabKey) {
    this.props.history.replace(`${this.props.match.path}/${tabKey}`);
  }

  onAppEdited(app) {
    if (app) {
      this.setState({
        apps: [app].concat(_.reject(this.state.apps, {name: app.name}))
      });
    }
  }

  onDataSourceEdited(dataSource) {
    if (dataSource) {
      this.setState({
        dataSources: [dataSource].concat(_.reject(this.state.dataSources, {name: dataSource.name}))
      });
    }
  }

  onAppDeleted(app) {
    if (app) {
      this.setState({
        apps: _.reject(this.state.apps, {name: app.name})
      });
    }
  }

  onDataSourceDeleted(dataSource) {
    if (dataSource) {
      this.setState({
        dataSources: _.reject(this.state.dataSources, {name: dataSource.name})
      });
    }
  }
}
