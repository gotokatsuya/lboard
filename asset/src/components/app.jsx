"use strict";

import React from "react";
import ReactDOM from "react-dom";
import { IndexRoute, Route } from "react-router";

import Home from "./view/home";
import Post from "./view/post";

var Routes = (
  <Route path="/" component={Home}>
    <IndexRoute component={Post}/>
    <Route name="post" path="post" component={Post}/>
  </Route>
);

export { Routes };
