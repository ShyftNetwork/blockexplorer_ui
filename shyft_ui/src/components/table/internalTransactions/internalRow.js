import React, { Component } from 'react';
import InternalTable from './internalTable';
import Pagination from '../pagination/pagination';
import classes from './table.css';
import axios from "axios";
import {API_URL} from "../../../constants/apiURL";
import ErrorMessage from './errorMessage';
import Loading from '../../UI materials/loading'

class InternalTransactionsTable extends Component {
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
        try {
            const response = await axios.get(`${API_URL}/get_internal_transactions_length/`);
            await this.setState({totalRecords: response.data});
            try {
                const response = await axios.get(`${API_URL}/get_internal_transactions/${currentPage}/${pageLimit}`);
                if(response.data === "\n") {
                    this.setState({emptyDataSet: true})
                } else {
                    this.setState({emptyDataSet: false})
                }
                await this.setState({data: response.data});
            } catch (err) {
                console.log(err);
            }
        } catch (err) {
            console.log(err);
        }
    }

    render() {
        let table;
        if(this.state.emptyDataSet === false && this.state.data.length > 0  ) {
            table = this.state.data.map((data, i) => {              
                return <InternalTable
                    key={`${data.tx_hash}${i}`}
                    Hash={data.tx_hash}
                    Action={data.internal_action}
                    To={data.to_address}
                    From= {data.from_address}
                    Gas={data.internal_gas}
                    GasUsed={data.gas_used}
                    ID={data.internal_id}
                    Input={data.internal_input}
                    Output={data.internal_output}
                    Time={data.internal_time}
                    Value={data.tx_amount}
                    detailInternalHandler={this.props.detailInternalHandler}            
                />
            });
       }

        let combinedClasses = ['responsive-table', classes.table];
        return (
            <div>     
                {
                    this.state.emptyDataSet === false && this.state.data.length > 0 ?  
                    <table className={combinedClasses.join(' ')}>
                        <thead>
                            <tr>                    
                                <th scope="col" className={classes.thItem}> Block Hash </th>
                                <th scope="col" className={classes.thItem}> Action </th>
                                <th scope="col" className={classes.thItem}> To </th>
                                <th scope="col" className={classes.thItem}> From </th>
                                <th scope="col" className={classes.thItem}> Gas </th>
                                <th scope="col" className={classes.thItem}> Gas Used</th>
                                <th scope="col" className={classes.thItem}> ID </th>
                                <th scope="col" className={classes.thItem}> Input </th>
                                <th scope="col" className={classes.thItem}> Output </th>
                                <th scope="col" className={classes.thItem}> Time </th>
                                <th scope="col" className={classes.thItem}> Value </th>
                            </tr>
                        </thead>
                        {table}
                        <div id={classes.pages}>
                                    <Pagination totalRecords={this.state.totalRecords} pageLimit={25} pageNeighbours={1} onPageChanged={this.onPageChanged} />
                        </div>
                    </table>
                    : <Loading>Internal Transactions</Loading>
                } 
            </div>
        );
    }
}
export default InternalTransactionsTable;
