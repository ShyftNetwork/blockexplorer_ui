import React, { Component } from 'react';
import AccountsTable from './accountsTable';
import classes from './accounts.css';
import axios from "axios/index";
import ErrorMessage from './errorMessage';
import {API_URL} from "../../../constants/apiURL";

class AccountTable extends Component {
    constructor(props) {
        super(props);
        this.state = {
            data: [],
            emptyDataSet: true
        };
    }

    async componentDidMount() {
        try {
            const response = await axios.get(`${API_URL}/get_all_accounts`);
            if(response.data === "\n") {
                this.setState({emptyDataSet: true})                                   
            } else {
                this.setState({emptyDataSet: false})                  
            }      
            await this.setState({data: response.data});
        } catch (err) {
            console.log(err);
        }
    }

    render() {
        let startNum = 1;
        let table;    
        if(this.state.emptyDataSet === false && this.state.data.length > 0  ) {
        const sorted = [...this.state.data];
        sorted.sort((a, b) => Number(a.Balance) > Number(b.Balance)); 
        table = sorted.reverse().map((data, i) => {
            const conversion = Number(data.Balance) / 10000000000000000000;
            const total = sorted
                .map(num => Number(num.Balance) / 10000000000000000000)
                .reduce((acc, cur) => acc + cur ,0);
            const percentage = ( (conversion / total) *100);
            return <AccountsTable
                key={`${data.addr}${i}`}
                Rank={startNum++}
                Percentage={percentage.toFixed(2)}
                Addr={data.Addr}
                Balance={conversion}
                AccountNonce={data.AccountNonce}
                detailAccountHandler={this.props.detailAccountHandler}
            />
        });
    }
          
        let combinedClasses = ['responsive-table', classes.table];
        return (      
            <div>     
                {
                   this.state.emptyDataSet === false && this.state.data.length > 0  ?  
                        <table className={combinedClasses.join(' ')}>
                            <thead>
                                <tr>
                                    <th scope="col" className={classes.thItem}>Rank</th>
                                    <th scope="col" className={classes.thItem}>Address</th>
                                    <th scope="col" className={classes.thItem}>Balance</th>
                                    <th scope="col" className={classes.thItem}>Percentage</th>
                                    <th scope="col" className={classes.thItem}>TxCount</th>
                                </tr>
                            </thead>
                            { table } 
                        </table>
                    : <ErrorMessage />
                } 
            </div>
        );
    }
}
export default AccountTable;
