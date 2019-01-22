import React, { Component } from 'react';

// import Pagination from '../pagination/pagination';
// import Loading from '../../UI materials/loading'
import {API_URL} from "../../../constants/apiURL";
import axios from "axios/index";
import TxTable from './table';

class Results extends Component {
    constructor(props) {
        super(props);
        this.state = {
            data: [],
            emptyDataSet: true
        };
    }

    async componentDidMount() {
        let pageLimit= 25;
        let currentPage = 1;

    }

    onPageChanged = async(data) => {
        const { currentPage, pageLimit } = data;

        try {
            const response = await axios.get(`${API_URL}/get_all_accounts/${currentPage}/${pageLimit}`);
            if(response.data === "\n") {
                this.setState({emptyDataSet: true})
            } else {
                this.setState({emptyDataSet: false})
            }
            await this.setState({data: response.data});
        } catch (err) {
            console.log(err);
        }
    };

    render() {
        let combinedClasses = ['responsive-table'];
        return (
            <div>
                <TxTable data={this.props.data}/>
            </div>
        );
    }
}
export default Results;
