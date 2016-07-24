"use strict";

import React from "react";
import Request from "request";
import PostList from "./post-list";

export default class PostAll extends React.Component {

  constructor(props) {
    super(props);
    this.state = { posts: null };
  }

  fetchPosts() {
  }

  componentDidMount() {
    this.fetchPosts();
  }

  render() {
    if (this.state.posts === null) {
      return (<div>Loading...</div>);
    }
    if (this.state.posts.length === 0) {
      return (<div></div>);
    }
    return (<PostList data={this.state.posts} />);
  }

}
