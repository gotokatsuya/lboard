"use strict";

import React from "react";
import PostAll from "./post-all";

export default class Post extends React.Component {

  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div>
        <h1>Posts</h1>
        <PostAll />
      </div>);
  }

}