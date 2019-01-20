import React, { Component } from 'react';
import MinedBlockTable from './blocksMinedTable';
import classes from './table.css';
import ErrorMessage from './errorMessage';

class BlocksMinedTable extends Component {
    constructor(props) {
        super(props);
        this.state = {
            data: []
        };
    }

    render() {
        const table = this.props.data.map((data, i) => {
            const conversion = data.block_rewards / 10000000000000000000;
            return <MinedBlockTable
                key={`${data.block_hash}${i}`}
                Hash={data.block_hash}
                Number={data.block_height}
                Coinbase={data.coinbase_hash}
                AgeGet={data.block_timestamp}
                GasUsed={data.block_gas}
                GasLimit={data.block_gaslimit}
                UncleCount={data.block_uncles}
                TxCount={data.block_txs}
                Reward={conversion}
                detailBlockHandler={this.props.detailBlockHandler}
                getBlocksMined={this.props.getBlocksMined}
            />
        });

        let combinedClasses = ['responsive-table', classes.table];
        return (
            <div>     
                {
                    this.props.data.length > 0 ?  
                        <table className={combinedClasses.join(' ')}>
                            <thead>
                                <tr>
                                    <th scope="col" className={classes.thItem}> Height </th>
                                    <th scope="col" className={classes.thItem}> Block Hash </th>
                                    <th scope="col" className={classes.thItem}> Age </th>
                                    <th scope="col" className={classes.thItem}> Txn </th>
                                    <th scope="col" className={classes.thItem}> Uncles </th>
                                    <th scope="col" className={classes.thItem}> Coinbase </th>
                                    <th scope="col" className={classes.thItem}> GasUsed </th>
                                    <th scope="col" className={classes.thItem}> GasLimit </th>
                                    <th scope="col" className={classes.thItem}> Avg.GasPrice </th>
                                    <th scope="col" className={classes.thItem}> Reward </th>
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
export default BlocksMinedTable;
