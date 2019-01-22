import React, { Component } from 'react';
import TransactionsTable from '../transactions/transactionTable'
import Pagination from '../pagination/pagination';
import classes from '../transactions/table.css';
import Loading from '../../UI materials/loading'

class TransactionTable extends Component {
    constructor(props) {
        super(props);
        this.state = {
            data: [],
            emptyDataSet: true
        };
    }

    render() {
        let table;
        if(this.props.data) {
            let t = [...this.state.data.push(this.props.data)];
            table = this.state.data.map((data, i) => {
                const conversion = data.tx_cost / 10000000000000000000;
                return <TransactionsTable
                    key={`${data.tx_hash}${i}`}
                    age={data.tx_timestamp}
                    txHash={data.tx_hash}
                    blockNumber={data.block_height}
                    to={data.to_address}
                    from={data.from_address}
                    value={data.tx_amount}
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
                   this.props.data ?
                        <table key={this.props.data.tx_hash} className={combinedClasses.join(' ')}>
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
                            {/*<div id={classes.pages}>*/}
                                {/*<Pagination totalRecords={25} pageLimit={25} pageNeighbours={1} onPageChanged={this.onPageChanged} />*/}
                            {/*</div>*/}
                        </table>
                        : <Loading>Transactions</Loading>
                }
            </div>
        );
    }
}
export default TransactionTable;
