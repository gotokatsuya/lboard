"use strict";

import gulp from "gulp";
import browserify from "browserify";
import babelify from "babelify";
import notify from "gulp-notify";
import glob from "glob";
import source from "vinyl-source-stream";

const entry = "./src/**/*.js";

gulp.task("scripts", () => {
    var src = glob.sync(entry);
    return browserify({
        debug: true,
        entries: src,
        extensions: [".jsx"]
    })
        .transform(babelify)
        .bundle()
        .on("error", notify.onError({ title: "scripts" }))
        .pipe(source("main.js"))
        .pipe(gulp.dest("../public"));
});

gulp.task("scripts:watch", () => {
    var src = glob.sync(entry);
    return gulp.watch(src, ["scripts"]);
});
