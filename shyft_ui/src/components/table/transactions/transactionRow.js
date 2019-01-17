import React, { Component } from 'react';
import TransactionsTable from './transactionTable';
import Pagination from '../pagination/pagination';
import classes from './table.css';
import axios from "axios";
import ErrorMessage from './errorMessage';
import {API_URL} from "../../../constants/apiURL";

class TransactionTable extends Component {
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
            const response = await axios.get(`${API_URL}/get_all_transactions_length`);
            await this.setState({totalRecords: response.data});
            try {
                const response = await axios.get(`${API_URL}/get_all_transactions/${currentPage}/${pageLimit}`);
                if(response.data === "\n") {
                    this.setState({emptyDataSet: true})
                } else {
                    this.setState({emptyDataSet: false})
                }
                console.log("data", response.data)
                await this.setState({data: response.data});
            } catch (err) {
                console.log(err);
            }
        } catch (err) {
            console.log(err);
        }
    }

    onPageChanged = async(data) => {
        const { currentPage, totalPages, pageLimit } = data;

        try {
            const response = await axios.get(`${API_URL}/get_all_transactions/${currentPage}/${pageLimit}`);
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
        let table;
        if(this.state.emptyDataSet === false && this.state.data.length > 0  ) {
            table = this.state.data.map((data, i) => {
                const conversion = data.Cost / 10000000000000000000;
                return <TransactionsTable
                    key={`${data.TxHash}${i}`}
                    age={data.Age}
                    txHash={data.TxHash}
                    blockNumber={data.BlockNumber}
                    to={data.ToGet}
                    from={data.From}
                    value={data.Amount}
                    cost={conversion}
                    getBlockTransactions={this.props.getBlockTransactions}
                    detailTransactionHandler={this.props.detailTransactionHandler}
                    detailAccountHandler={this.props.detailAccountHandler}
                />
            })        
        }

        let combinedClasses = ['responsive-table', classes.table];
        return (
            <div>     
                {
                    this.state.emptyDataSet === false && this.state.data.length > 0 ?  
                        <table key={this.state.data.TxHash} className={combinedClasses.join(' ')}>
                            <thead>
                                <tr>
                                    <th scope="col" className={classes.thItem}> TxHash </th>
                                    <th scope="col" className={classes.thItem}> Block </th>
                                    <th scope="col" className={classes.thItem}> Age </th>
                                    <th scope="col" className={classes.thItem}> From </th>                      
                                    <th scope="col" className={classes.thItem}> To </th>
                                    <th scope="col" className={classes.thItem}> Value </th>
                                    <th scope="col" className={classes.thItem}> TxFee </th>
                                </tr>
                            </thead>
                            {table}
                            <div id={classes.pages}>
                                <Pagination totalRecords={this.state.totalRecords} pageLimit={25} pageNeighbours={1} onPageChanged={this.onPageChanged} />
                            </div>
                        </table>
                    : <ErrorMessage />
                } 
            </div>           
        );
    }
}
export default TransactionTable;
