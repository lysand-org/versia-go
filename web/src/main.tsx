import "./instrument";

import React from "react";
import * as Sentry from "@sentry/react";
import ReactDOM from "react-dom/client";
import {ReactQueryDevtools} from "@tanstack/react-query-devtools";
import "./index.css";
import {RouterProvider} from "@tanstack/react-router";
import {QueryClient, QueryClientProvider} from "@tanstack/react-query";
import {router} from "./routing.ts";

const queryClient = new QueryClient();

const App = Sentry.withProfiler(() => (
    <React.StrictMode>
        <QueryClientProvider client={queryClient}>
            <RouterProvider router={router}/>
            <ReactQueryDevtools/>
        </QueryClientProvider>
    </React.StrictMode>
));

ReactDOM.createRoot(document.getElementById("root")!).render(<App/>);
