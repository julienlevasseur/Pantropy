import React, { Component } from 'react';
import SideNav, { Nav, NavIcon, NavText } from 'react-sidenav';
import SvgIcon from 'react-icons-kit';
//import logo from './logo.svg';
import './App.css';
//import '../node_modules/bootstrap/dist/css/bootstrap.css';

import { ic_aspect_ratio } from 'react-icons-kit/md/ic_aspect_ratio';
import { ic_cloud } from 'react-icons-kit/md/ic_cloud';
import { cog } from 'react-icons-kit/icomoon/cog';

class App extends Component {

  render() {
    return (
      <div className="App">
        <header className="App-header">
          <h1 className="App-title" align="left">Pantropy</h1>
        </header>
        <div className="App-menu">
          <SideNav highlightColor='#fff' highlightBgColor='#00bcd4' defaultSelected='sales'>       
            <Nav id='dashboard'>
                <NavIcon><SvgIcon size={20} icon={ic_aspect_ratio}/></NavIcon>    
                <NavText> Dashboard </NavText>
            </Nav>
            <Nav id='providers'>
                <NavIcon><SvgIcon size={20} icon={ic_cloud}/></NavIcon>
                <NavText> Providers </NavText>
            </Nav>
            <Nav id='resources'>
                <NavIcon><SvgIcon size={20} icon={cog}/></NavIcon>
                <NavText> Resources </NavText>
            </Nav>
          </SideNav>
        </div>
        <div className="App-main-panel" background="#ccc">
          pouet
        </div>
      </div>
    );
  }
}

export default App;
