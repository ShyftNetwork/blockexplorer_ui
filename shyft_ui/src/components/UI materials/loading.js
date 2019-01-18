import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import CircularProgress from '@material-ui/core/CircularProgress';
import { Typography } from '@material-ui/core';

const styles = theme => ({
  progress: {
    margin: theme.spacing.unit * 2,
    
  },
  position: {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    justifyContent: 'center'
  },
});

function CircularIndeterminate(props) {
  const { classes } = props;
  return (
        <div className={classes.position}>
            <CircularProgress className={classes.progress} size={120} thickness={5} />
            <Typography>Loading {props.children}...</Typography>
        </div>
  );
}

CircularIndeterminate.propTypes = {
  classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(CircularIndeterminate);