"use strict";

import React from "react";
import { NavItem } from "react-bootstrap";
import { LinkContainer } from "react-router-bootstrap";

export default class Home extends React.Component {

    constructor(props) {
        super(props);
        this.state = { me: null };
    }

    fetchMe() {
    }

    componentDidMount() {
        this.fetchMe();
    }

    handleRefreshMe(me) {
    }

    render() {
        if (this.state.me === null) {
            return (<div>Loading</div>)
        }
        return (
            <div>
                <div>
                    <div>board</div>
                    <div>
                        <LinkContainer to="/post">
                            <NavItem>All Posts</NavItem>
                        </LinkContainer>
                    </div>
                </div>
                {React.cloneElement(this.props.children, { onRefreshMe: this.handleRefreshMe.bind(this) }) }
            </div>
        );
    }
}
