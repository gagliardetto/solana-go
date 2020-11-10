/**
 * Copyright 2019 dfuse Platform Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

function refreshToken(server, token, apiKey, onCompletion, onError) {
    if(token === "") {
        return getToken(apiKey, onCompletion, onError)
    }

    const jwt = parseJwt(token);
    const exp = jwt["exp"];
    const now = Date.now() / 1000;

    console.log("exp  : " + exp);
    console.log("now  : " + now);
    const remainingTime = exp - now;

    console.log("Time remaining in second: " + remainingTime);
    if (remainingTime < 60 * 60) {
        return getToken(apiKey, onCompletion, onError)
    }

    onCompletion(token);
}

function getToken(apiKey, onCompletion, onError) {
    const url = "https://auth.dfuse.io/v1/auth/issue";
    const r = new XMLHttpRequest();
    r.open("POST", url, false);
    r.setRequestHeader("Content-type", "application/json");
    r.onreadystatechange = function() {//Call a function when the state changes.
        if(r.readyState === 4) {
            if (r.status === 200) {
                console.log("Got new token: " + r.response);
                const responseToken =  JSON.parse(r.response)
                onCompletion(responseToken["token"]);
            } else {
                alert("Error: " + r.status + " - " + r.response);
                onError(r.status, r.response);
            }
        }
    };

    r.send('{"api_key":"' + apiKey + '"}');
}

function graphQLFetcherFactory(urlQueryParams) {
    let graphqlUrl = "/graphql"
    if (isAlphaSchemaQueryParamFound(urlQueryParams)) {
        console.log("Requesting alpha endpoints")
        graphqlUrl += "?alpha-schema=true"
    }

    return (graphQLParams) => {
        return fetch(graphqlUrl, {
            method: "post",
//            headers: {
//                Authorization: "Bearer " + token,
//            },
            body: JSON.stringify(graphQLParams),
            credentials: "include",
        }).then(function (response) {
            return response.text();
        }).then(function (responseBody) {
            try {
                return JSON.parse(responseBody);
            } catch (error) {
                return error.message;
            }
        }).catch(function (error) {
            console.log("Error:", error);
            alert(error)
        });
    }
}

function isAlphaSchemaQueryParamFound(urlQueryParams) {
    if (urlQueryParams.has("alpha-schema")) {
        return urlQueryParams.get("alpha-schema") === "true"
    }

    if (urlQueryParams.has("alphaSchema")) {
        return urlQueryParams.get("alphaSchema") === "true"
    }

    return false
}

function fetchQueryProp(queryParams) {
    if (!queryParams.has("query")) {
        return undefined
    }

    try {
        return window.atob(queryParams.get("query"))
    } catch (error) {
        console.error("query params 'query' is not a valid base64 object")
        return undefined
    }
}

function fetchVariablesProp(queryParams) {
    if (!queryParams.has("variables")) {
        return undefined
    }

    try {
        return window.atob(queryParams.get("variables"))
    } catch (error) {
        console.error("query params 'variables' is not a valid base64 object")
        return undefined
    }
}

function pushState(url, query, variables) {
    const queryParams = []
    if (query !== undefined) {
        queryParams.push(`query=${window.btoa(query)}`)
    }

    if (variables !== undefined) {
        queryParams.push(`variables=${window.btoa(variables)}`)
    }

    if (queryParams.length <= 0) {
        return
    }

    window.history.pushState("", "New Query", `${url.pathname}?${queryParams.join("&")}`);
}

async function getConfig() {
    if (window.location.hostname === "localhost") {
        return await fetchConfig()
    }

    const parts = window.location.host.split(".");
    return { network: parts[0], protocol: parts[1] }
}

function fetchFavorites() {
    console.info("Fetching favorites JSON data")
    return fetch("/graphiql/favorites.json")
            .then((response) => response.json())
            .then((body) => {
                console.log("Got favorites JSON data")
                return body
            })
            .catch((error) => {
                console.log("Fetch favorites JSON data error", error);
                return {}
            })
}

function fetchConfig() {
    console.info("Fetching config JSON data")
    return fetch("/graphiql/config.json")
            .then((response) => response.json())
            .then((body) => {
                console.log("Got config JSON data")
                return body
            })
            .catch((error) => {
                console.log("Fetch config JSON data error", error);
                return {}
            })
}

function getFavoriteFromStorage() {
    const storageItem = localStorage.getItem("graphiql:favorites");
    console.log("Retrieved client favorites from browser storage")

    store = { favorites: [] };
    if (storageItem !== null) {
        store = JSON.parse(storageItem);
    }

    return store
}

function setFavoriteFromStorage(store) {
    console.log("Saving client favorites to browser storage")
    localStorage.setItem("graphiql:favorites", toJSON(store));
}

async function reconfigureGraphiQLStorage(protocol, network, alphaSchema) {
    const isFirstTime = localStorage.getItem("dfuse:graphiql:is_first_time");
    if (isFirstTime == null) {
        // Let's open the history pane the first time the user opens this page
        localStorage.setItem("graphiql:historyPaneOpen", toJSON(true));
    }

    const serverFavorites = await setFavorites(protocol, network, alphaSchema)

    localStorage.setItem("dfuse:graphiql:is_first_time", toJSON(false))

    if (isFirstTime == null && serverFavorites.length > 0) {
        localStorage.setItem("graphiql:query", serverFavorites[0].query);

        if (serverFavorites[0].variables) {
            localStorage.setItem("graphiql:variables", serverFavorites[0].variables);
        }
    }
}

async function setFavorites(protocol, network, alphaSchema) {
    const favoritesByProtocolMap = await fetchFavorites()

    console.log(`Looking for favorites for given ${protocol}/${network} values`)
    const serverFavorites = favoritesByProtocolMap[protocol]
    if (serverFavorites == null) {
        console.log("Favorites not found for this protocol/network values.")
        return
    }

    const store = getFavoriteFromStorage()

    // Clear all dfuse managed favorites, we will add them back
    store["favorites"] = store["favorites"].filter((value) =>
        // We keep only favorites that don't have the `procotol`
        value.protocol == null
    )

    console.log("Favorites store prior update")
    serverFavorites.reverse().forEach((favorite) => {
        if (!alphaSchema && favorite.alpha) {
            return
        }

        favorite.favorite = true
        favorite.protocol = protocol
        if (favorite.variables && typeof favorite.variables === "object") {
            const networkVariables = favorite.variables[network]
            if (networkVariables != null) {
                favorite.variables = networkVariables
            }
        }

        store["favorites"] = updateFavorite(store["favorites"], favorite);
    })
    console.log("Favorites store after update")

    setFavoriteFromStorage(store)

    // We reverse it again because the `reverse` operation is "in-place"
    return serverFavorites.reverse()
}

function updateFavorite(favorites, fav) {
    const index = favorites.findIndex(f => (f.label === fav.label));
    if (index >= 0) {
        console.log(`Updating favorite ${fav.label}`)
        favorites[index] = fav
    } else {
        console.log(`Adding favorite ${fav.label}`)
        favorites.push(fav)
    }

    return favorites
}

function toJSON(input) {
    return JSON.stringify(input)
}
