import { useRecordContext } from "ra-core";
import React from "react";
import { makeStyles } from "@material-ui/core";
import LaunchIcon from "@material-ui/icons/Launch";

const useStyles = makeStyles({
  link: {
    textDecoration: "none",
  },
  icon: {
    width: "0.5em",
    height: "0.5em",
    paddingLeft: 2,
  },
});

export const MyUrlField = ({ source }) => {
  const record = useRecordContext();
  const classes = useStyles();
  return record ? (
    <a href={record[source]} className={classes.link}>
      {record[source]}
      <LaunchIcon className={classes.icon}></LaunchIcon>
    </a>
  ) : null;
};
