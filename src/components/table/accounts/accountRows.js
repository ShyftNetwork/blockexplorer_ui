import React, { Component } from 'react';
import AccountsTable from './accountsTable';
import Pagination from '../pagination/pagination';
import classes from './accounts.css';
import axios from "axios/index";
import ErrorMessage from './errorMessage';
import Loading from '../../UI materials/loading'
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
        let pageLimit= 25;
        let currentPage = 1;
        try {
            const response = await axios.get(`${API_URL}/get_all_accounts_length`);
            await this.setState({totalRecords: response.data});
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
        } catch (err) {
            console.log(err);
        }
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
        let startNum = 1;
        let table;    
        if(this.state.emptyDataSet === false && this.state.data.length > 0  ) {
        const sorted = [...this.state.data];
        sorted.sort((a, b) => Number(a.balance) > Number(b.balance));
        table = sorted.reverse().map((data, i) => {
            const conversion = Number(data.balance) / 10000000000000000000;
            const total = sorted
                .map(num => Number(num.balance) / 10000000000000000000)
                .reduce((acc, cur) => acc + cur ,0);
            const percentage = ( (conversion / total) *100);
            const Perc=percentage.toFixed(2);
	    var PercUpdated='';
            if(Perc.includes('NaN') === true){
 		//console.log(Perc);
                //Perc.replace(/NaN/, '0');
		PercUpdated = Perc.replace(/NaN/, '0') ;
 		//console.log(Perc.replace(/NaN/, '0'));
 		//console.log(PercUpdated);
 		//console.log('Found NaN After ');
	    }else{
 		console.log('Did not Find NaN');
	    }
            return <AccountsTable
                key={`${data.address}${i}`}
                Rank={startNum++}
                Percentage={PercUpdated}
                Addr={data.address}
                Balance={conversion}
                AccountNonce={data.nonce}
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
                            <div id={classes.pages}>
                                <Pagination totalRecords={this.state.totalRecords.page_count} pageLimit={25} pageNeighbours={1} onPageChanged={this.onPageChanged} />
                            </div>
                        </table>
                    : <Loading>Accounts</Loading>
                } 
            </div>
        );
    }
}
export default AccountTable;
