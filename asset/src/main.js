"use strict";

import React from "react";
import { render } from "react-dom";
import { Router,hashHistory } from "react-router";
import { Routes } from "./components/app"

window.onload = function () {
  render(
    <Router history={hashHistory}>{Routes}</Router>,
    document.getElementById("root")
  );
}
