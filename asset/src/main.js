"use strict";

import React from "react";
import { render } from "react-dom";
import { Router } from "react-router";

import { Routes } from "./components/app"

window.onload = function () {
  render(
    <Router>{Routes}</Router>,
    document.getElementById("root")
  );
}
