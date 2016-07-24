"use strict";

import React from "react";
import { NavItem } from "react-bootstrap";
import { Link } from "react-router";

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
        return (
            <div>
                <div>
                    <div>board</div>
                    <div>
                        <Link to="/post">
                            Posts
                        </Link>
                    </div>
                </div>
                {React.cloneElement(this.props.children, { onRefreshMe: this.handleRefreshMe.bind(this) }) }
            </div>
        );
    }
}
