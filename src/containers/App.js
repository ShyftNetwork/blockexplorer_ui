import React, { Component } from "react";
import axios from 'axios';
import Nav from "../components/nav/nav";
import { BrowserRouter, Route  } from 'react-router-dom'
import { API_URL } from "../constants/apiURL";
import Modal from 'react-bootstrap/lib/Modal';
import Button from 'react-bootstrap/lib/Button';
import Grid from 'react-bootstrap/lib/Grid';
import Col from 'react-bootstrap/lib/Col';
import Row from 'react-bootstrap/lib/Row';

///**LANDING PAGE**///
import Home from '../components/home/home';

///**TRANSACTIONS**///
import TransactionRow from '../components/table/transactions/transactionRow';
import DetailTransactionTable from "../components/table/transactions/transactionDetailsRow";
import BlockTxs from "../components/table/transactions/blockTx";


///**INTERNAL TRANSACTIONS**///
import InternalTransactionRow from '../components/table/internalTransactions/internalRow';

///**BLOCKS**///
import BlocksRow from '../components/table/blocks/blockRows';
import DetailBlockTable from '../components/table/blocks/blocksDetailsRow';
import BlockDetailHeader from "../components/nav/blockHeaders/blockDetailHeader";
import BlocksMinedTable from "../components/table/blocks/blocksMined";

///**ACCOUNTS**///
import AccountsRow from '../components/table/accounts/accountRows';
import DetailAccountsTable from "../components/table/accounts/detailAccountsRow";

//**SEARCH QUERY**//
import QueryResults from '../components/table/searchResults/results';

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
        blockDetailData: [],
        transactionDetailData: [],
        accountDetailData: [],
        internalDetailData: [],
        blocksMined: [],
        blockTransactions: [],
        accounts: [],
        reqAccount: '',
        overlayTriggered: false,
        overlayContent: "",
        data: []
    };
  }

    detailBlockHandler = async(blockNumber) => {
        try {
            const response = await axios.get(`${API_URL}/get_block/${blockNumber}`);
            await this.setState({ blockDetailData: response.data, overlayTriggered: true, overlayContent: "block" })
        }
        catch(error) {
           console.log(error)
        }
    };

    detailTransactionHandler = async(txHash) => {
        try {
            const response = await axios.get(`${API_URL}/get_transaction/${txHash}`);
            await this.setState({ transactionDetailData: response.data, overlayTriggered: true, overlayContent: "transaction" })
        }
        catch(error) {
            console.log(error)
        }
    };


    detailAccountHandler = async(addr) => {
        let pageLimit= 25;
        let currentPage = 1;
        try {
            const response = await axios.get(`${API_URL}/get_account_txs/${currentPage}/${pageLimit}/${addr}`);
            await this.setState({ accountDetailData: response.data, reqAccount: addr, overlayTriggered: false, overlayContent: "account"  })
        }
        catch(error) {
            console.log(error)
        }
    };

    detailInternalHandler = async(txHash) => {
        let pageLimit= 25;
        let currentPage = 1;
        try {
            const response = await axios.get(`${API_URL}/get_internal_transactions/${currentPage}/${pageLimit}/${txHash}`);
            await this.setState({ internalDetailData: response.data, reqAccount: txHash, overlayTriggered: true, overlayContent: "internal" })
        }
        catch(error) {
            console.log(error)
        }
    };

    searchQueryHandler = async(query) => {
        let pageLimit= 25;
        let currentPage = 1;
        try {
            const response = await axios.get(`${API_URL}/search/${query}`);
            await this.setState({data: response.data})
            return response.data
            //await this.setState({ data: response.data });
        }
        catch(error) {
            console.log(error)
        }
    };


    getInternalTransactions = async() => {
        let pageLimit= 25;
        let currentPage = 1;
        try {
            const response = await axios.get(`${API_URL}/get_internal_transactions/${currentPage}/${pageLimit}`);
            await this.setState({ internalTransactions: response.data,  overlayTriggered: false })
        }
        catch(error) {
            console.log(error)
        }
    };

    getBlockTransactions = async(blockNumber) => {
        let pageLimit= 25;
        let currentPage = 1;
        try {
            const response = await axios.get(`${API_URL}/get_all_transactions_from_block/${currentPage}/${pageLimit}/${blockNumber}`);
            await this.setState({ blockTransactions: response.data, reqBlockNum: blockNumber })
        }
        catch(error) {
            console.log(error)
        }
    };

    getBlocksMined = async(coinbase) => {
        let pageLimit= 25;
        let currentPage = 1;
        try {
            const response = await axios.get(`${API_URL}/get_blocks_mined/${currentPage}/${pageLimit}/${coinbase}`);
            await this.setState({ blocksMined: response.data, reqCoinbase: coinbase })
        }
        catch(error) {
            console.log(error)
        }
    };

    hideOverlay = () => {
        this.setState({overlayTriggered: false})
    };

    getBlockData = (data, page) => {
        let components = [];
        let dataEntry;
        if(page === "account" || page === "internal") {
            dataEntry = data[0];
        }
        else { 
            dataEntry = data;
        }
        for (let key in dataEntry) {
            let value = dataEntry[key];
            if (dataEntry["To"] === null) {
                delete dataEntry["To"]
            }
            if( key === "Input" || key === "Output" || key === "Data") {
                components.push( 
                    <Grid>
                        <Row className="show-grid">
                            <Col xs={3}  md={3}  style={{fontSize:'6.5pt', color: '#565656', paddingTop: '10pt', fontFamily: 'Open Sans, sans-serif'}}>
                                {key}
                            </Col>
                            <Col xs={9}  md={9}  style={{fontSize:'6.5pt', color: '#565656', paddingTop: '5pt', marginLeft:'-5pt', fontFamily: 'Open Sans, sans-serif'}}>
                                <Button style={{color: '#8f67c9'}} bsStyle="link" bsSize="small" onClick={()=> alert( value )}> Show {key} </Button> 
                            </Col>                          
                        </Row>
                    </Grid>
                );             
            } else {
                components.push( 
                    <Grid>
                        <Row className="show-grid">
                            <Col xs={3}  md={3}  style={{fontSize:'6.5pt', color: '#565656', paddingTop: '10pt', fontFamily: 'Open Sans, sans-serif'}}>
                                {key}
                            </Col>
                            <Col xs={9}  md={9}  style={{fontSize:'6.5pt', color: '#565656', paddingTop: '10pt', fontFamily: 'Open Sans, sans-serif'}}>
                                {value}
                            </Col>
                        </Row>
                    </Grid>
                ); 
            }
        }
        return components;
    }

    renderOverlay = () => {
        const page = this.state.overlayContent;
        let data, title;
        switch (page) {
            case "block" : 
                title = "BLOCK OVERVIEW";
                data = this.state.blockDetailData;
            break; 
            case "transaction" :
                title = "TRANSACTION OVERVIEW";
                data =  this.state.transactionDetailData;
            break;
            case "account" : 
                title = "ACCOUNT OVERVIEW";
                data =  this.state.accountDetailData;
            break;
            case "internal" : 
                title = "Internal Transaction Overview";
                data = this.state.internalDetailData;
            break;
            default: console.log("error");
        }
        return ( 
            <div className="static-modal">
                <Modal.Dialog>
                    <Modal.Header>
                        <Modal.Title style={{color: '#593c83', fontSize: '8pt', letterSpacing: '2pt',  fontFamily: 'Open Sans, sans-serif' }}> {title} </Modal.Title>
                    </Modal.Header>
                    <Modal.Body>
                        <ul>
                        {this.getBlockData(data, page)}         
                        </ul>       
                    </Modal.Body>
                    <Modal.Footer>
                    <Button onClick={this.hideOverlay}>Close</Button>                   
                    </Modal.Footer>
                </Modal.Dialog>
            </div>
        );
    };

  render() {
    return (
        <BrowserRouter>
            <div>
                
                {
                    this.state.overlayTriggered ? this.renderOverlay() : null
                }
        
                <div style={ this.state.overlayTriggered ? {backgroundColor:"#f7f8f9", paddingBottom:"5%", opacity:0.5, zIndex: -10000, height: '100%', width: '100%'}  : {backgroundColor:"#f7f8f9", paddingBottom:"5%", height: '100%', width: '100%' }  }>
                    <Nav searchQueryHandler={this.searchQueryHandler}/>
        
                    <Route path="/" exact render={({ match }) =>
                        <Home style={{width: '100%'}}/> 
                        }
                    />

                    <Route path="/transactions" render={({match}) =>
                        <div>                  
                            <TransactionRow
                                getBlockTransactions={this.getBlockTransactions}
                                detailTransactionHandler={this.detailTransactionHandler}
                                detailAccountHandler={this.detailAccountHandler}/>
                        </div>
                        }
                    />

                    <Route path="/blocks" exact render={({ match }) =>
                        <div>                        
                            <BlocksRow
                                getBlocksMined={this.getBlocksMined}
                                getBlockTransactions={this.getBlockTransactions}
                                detailBlockHandler={this.detailBlockHandler}/>                        
                        </div>
                        }
                    />

                    <Route path="/internalTransactions" exact render={({match}) =>
                        <div>                        
                            <InternalTransactionRow                              
                                getInternalTransactionsHandler={this.getInternalTransactions}
                                detailInternalHandler={this.detailInternalHandler}
                                />                        
                        </div>
                        }
                    />

                    <Route path="/accounts" exact render={({match}) =>
                        <div>         
                            <AccountsRow detailAccountHandler={this.detailAccountHandler}/>
                        </div>
                        }
                    />

                    <Route path="/transaction/details" exact render={({match}) =>
                    <div>
                        <DetailTransactionTable
                            data={this.state.transactionDetailData}/>
                    </div>}
                    />

                    <Route path="/blocks/detail" exact render={({match}) =>
                    <div>
                        <BlockDetailHeader
                            blockNumber={this.state.blockDetailData.Number}/>
                        <DetailBlockTable
                            getBlockTransactions={this.getBlockTransactions}
                            data={this.state.blockDetailData}/>
                    </div>}
                    />

                    <Route path="/block/transactions" exact render={({match}) =>
                    <div>
                        <BlockTxs
                            data={this.state.blockTransactions}
                            getBlockTransactions={this.getBlockTransactions}
                            detailTransactionHandler={this.detailTransactionHandler}
                            detailAccountHandler={this.detailAccountHandler}/>

                    </div>}
                    />

                    <Route path="/mined/blocks" exact render={({match}) =>
                    <div>                    
                        <BlocksMinedTable
                            getBlockTransactions={this.getBlockTransactions}
                            getBlocksMined={this.getBlocksMined}
                            data={this.state.blocksMined}/>
                    </div>}
                    />

                    <Route path="/account/detail" exact render={({match}) =>
                    <div>                     
                        <DetailAccountsTable
                            transactionDetailHandler={this.detailTransactionHandler}
                            addr={this.state.reqAccount}
                            data={this.state.accountDetailData}/>
                    </div>}
                    />
                    <Route path="/search/" exact render={({match}) =>
                        <div>
                            <QueryResults
                                //transactionDetailHandler={this.detailTransactionHandler}
                                data={this.state.data}/>
                        </div>}
                    />
                </div>
            </div>
        </BrowserRouter>
    );
  }
}

export default App;
