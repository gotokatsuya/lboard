"use strict";

import gulp from "gulp";

gulp.task("watch", [
    "eslint:watch",
    "scripts:watch"
]);
