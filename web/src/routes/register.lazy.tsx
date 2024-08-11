import {createLazyFileRoute} from "@tanstack/react-router";
import {useCreateUser} from "../hooks/useCreateUser.ts";
import {UserCreation} from "../api/auth.ts";

export const Route = createLazyFileRoute("/register")({
    component: Register,
});

function Register() {
    const register = useCreateUser();

    return (
        <div className="p-2">
            <button
                onClick={() => register.mutate(new UserCreation({username: "user"}))}
                disabled={register.status === "pending"}
            >
                Register
            </button>
        </div>
    );
}
