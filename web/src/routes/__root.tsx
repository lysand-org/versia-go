import {createRootRoute, Link, Outlet} from "@tanstack/react-router";
import {TanStackRouterDevtools} from "@tanstack/router-devtools";
import {ReactNode} from "react";

function Wrapper({children}: { children: ReactNode }) {
    return (
        <>
            <div className="p-2 flex gap-2">
                <Link to="/" className="[&.active]:font-bold">
                    Home
                </Link>
            </div>
            <hr/>
            {children}
            <TanStackRouterDevtools/>
        </>
    );
}

export const Route = createRootRoute({
    component: () => (
        <Wrapper>
            <Outlet/>
        </Wrapper>
    ),
    // errorComponent: (props) => (
    //   <Wrapper>
    //     <ErrorHandler {...props} />
    //   </Wrapper>
    // ),
});
