import React from "react"
import App from "./App"
import ApolloClient from "apollo-client";
import {InMemoryCache} from "apollo-cache-inmemory";
import {createHttpLink} from "apollo-link-http";
import {ApolloProvider} from "@apollo/react-hooks"

const baseURL = process.env.REACT_APP_BASE_URL || "http://localhost:8080"
const httpLink = createHttpLink({
    uri: `${baseURL}/query`
})

const client = new ApolloClient({
    link: httpLink,
    cache: new InMemoryCache()
})

const Provider = () => {
    return (
        <ApolloProvider client={client}>
            <App/>
        </ApolloProvider>
    )
}

export default Provider