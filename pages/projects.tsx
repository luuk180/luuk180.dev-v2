import NavBar from "@/components/navbar";
import { ReactElement, JSXElementConstructor, ReactFragment, ReactPortal } from "react";
import {object} from "prop-types";

export async function getServerSideProps() {
    return {
        props: {entries: []},
    }
}

function Projects({entries}) {
    return <>
        <NavBar/>
        <div className="container mx-auto py-4 w-full text-white text-l">
            <p>
                This will be my projects!
            </p>
            <ul>
                {entries.map((entry) => <li>{entry.name}</li>)}
            </ul>
        </div>
    </>
}

export default Projects;