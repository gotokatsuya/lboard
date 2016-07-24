"use strict";

import React from "react";
import PostCell from "./post-cell";

export default class PostList extends React.Component {

  constructor(props) {
    super(props);
  }

  render() {
    let posts = this.props.data.map(function (row, index) {
      return (<PostCell key={index} data={row}/>);
    });
    return (<div>{posts}</div>);
  }

}
