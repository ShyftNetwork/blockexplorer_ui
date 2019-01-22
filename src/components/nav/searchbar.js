import React, { Component } from "react";
import PropTypes from 'prop-types';
import { Redirect } from 'react-router-dom'
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import InputBase from '@material-ui/core/InputBase';
import { fade } from '@material-ui/core/styles/colorManipulator';
import { withStyles } from '@material-ui/core/styles';
import SearchIcon from '@material-ui/icons/Search';

//class App extends Component
class SearchAppBar extends Component {
    constructor(props) {
        super(props);
        this.state = {
            change: "",
            data: [],
            results: false,
        };
    }

    handleChange = (e) => {
        let change = e.target.value;
        this.setState({value:change})
    };

    keyPress = async(e) =>{
        //
        if(e.keyCode === 13){
            try {
                let res = await this.props.searchQueryHandler(this.state.value)
                this.setState({data: [...this.state.data.push(res)], results: true})
            }
            catch(error) {
                console.log(error)
            }
        }
    };

    render() {
        const {classes} = this.props;
        return (
            <div className={classes.root}>
                <AppBar className={classes.bar} position="static">
                    <Toolbar>
                        <Typography className={classes.title} variant="h6" color="inherit" noWrap>
                            Block Explorer
                        </Typography>
                        <div className={classes.grow}/>
                        <div className={classes.search}>
                            <div className={classes.searchIcon}>
                                <SearchIcon/>
                            </div>
                            <InputBase
                                onChange={this.handleChange}
                                onKeyDown={this.keyPress}
                                placeholder="Search…"
                                classes={{
                                    root: classes.inputRoot,
                                    input: classes.inputInput,
                                }}
                            />
                        </div>
                    </Toolbar>
                </AppBar>
                {this.state.results ?
                <Redirect to={{
                    pathname: '/search/',
                    state: {data: this.state.data}
                }}/> : null
                }
            </div>
        );
    }
}

const styles = theme => ({
    root: {
        width: '97%',
    },
    bar: {
        backgroundColor: '#4f2e7e',
        boxShadow: 'none',
    },
    grow: {
        flexGrow: 1,
    },
    menuButton: {
        marginLeft: -12,
        marginRight: 20,
    },
    title: {
        display: 'none',
        [theme.breakpoints.up('sm')]: {
            display: 'block',
        },
        marginLeft: '-20px'
    },
    search: {
        position: 'relative',
        borderRadius: theme.shape.borderRadius,
        backgroundColor: fade(theme.palette.common.white, 0.15),
        '&:hover': {
            backgroundColor: fade(theme.palette.common.white, 0.25),
        },
        marginLeft: 0,
        width: '100%',
        [theme.breakpoints.up('sm')]: {
            marginLeft: theme.spacing.unit,
            width: 'auto',
        },
    },
    searchIcon: {
        width: theme.spacing.unit * 9,
        height: '100%',
        position: 'absolute',
        pointerEvents: 'none',
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
    },
    inputRoot: {
        color: 'inherit',
        width: '100%',
    },
    inputInput: {
        paddingTop: theme.spacing.unit,
        paddingRight: theme.spacing.unit,
        paddingBottom: theme.spacing.unit,
        paddingLeft: theme.spacing.unit * 10,
        transition: theme.transitions.create('width'),
        width: '100%',
        [theme.breakpoints.up('sm')]: {
            width: 120,
            '&:focus': {
                width: 200,
            },
        },
    },
});


SearchAppBar.propTypes = {
  classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(SearchAppBar);