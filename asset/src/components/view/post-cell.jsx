"use strict";

import React from "react";

export default class PostCell extends React.Component {

  constructor(props) {
    super(props);
  }

  render() {
    let post = this.props.data;
    if (!post) {
      return <div></div>
    }
    return (<div></div>);
  }

}
