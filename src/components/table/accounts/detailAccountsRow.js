import React, { Component } from 'react';
import DetailAccountsTable from './detailAccountsTable';
import ErrorMessage from "./errorMessage";
import classes from './table.css';

class AccountTransactionTable extends Component {
    render() {
        let table;
        if(this.props.data.length < 1) {
           return <ErrorMessage />
        } else {
            table = this.props.data.map((data, i) => {
                const costConversion = data.tx_cost / 10000000000000000000;
                const amountConversion = data.tx_amount / 10000000000000000000;
                return <DetailAccountsTable
                    key={`${data.tx_hash}${i}`}
                    age={data.tx_timestamp}
                    txHash={data.tx_hash}
                    blockNumber={data.block_height}
                    to={data.to_address}
                    from={data.from_address}
                    value={amountConversion}
                    cost={costConversion}
                    addr={this.props.addr}
                    detailTransactionHandler={this.props.transactionDetailHandler}
                />
            })
        }

        let combinedClasses = ['responsive-table', classes.table];
        return (
            <div>
                {
                    this.props.data.length > 0 ?
                    <table key={this.props.data.TxHash} className={combinedClasses.join(' ')}>
                        <thead className={classes.tHead}>
                        <tr>
                            <th scope="col" className={classes.thItem}>TxHash</th>
                            <th scope="col" className={classes.thItem}>Block</th>
                            <th scope="col" className={classes.thItem}>Age</th>
                            <th scope="col" className={classes.thItem}>From</th>
                            <th scope="col" className={classes.thItem}> </th>
                            <th scope="col" className={classes.thItem}>To</th>
                            <th scope="col" className={classes.thItem}>Value</th>
                            <th scope="col" className={classes.thItem}>TxFee</th>
                        </tr>
                        </thead>
                        {table}
                    </table>
                    : <ErrorMessage />
                }
            </div>
        );
    }
}
export default AccountTransactionTable;
