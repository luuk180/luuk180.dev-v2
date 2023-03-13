import NavBar from "@/components/navbar";
import {GetServerSideProps} from "next";

export const getServerSideProps: GetServerSideProps = async () => {
    const response = await fetch('https://api.github.com/graphql', {
        method: 'POST',
        headers: {
            Authorization: `bearer ${process.env.GITHUB_BEARER}`,
        },
        body: JSON.stringify({
            query: `
               {
                 user(login: "luuk180") {
                   repositories(orderBy: {field: PUSHED_AT, direction: ASC}, first: 100) {
                     edges {
                       node {
                         isPrivate
                         name
                         url
                         homepageUrl
                         description
                         diskUsage
                       }
                     }
                   }
                 }
               }
               `,
        }),
    });
    const json = await response.json();
    let entries = json.data.user.repositories.edges;
    if (entries === undefined) {
        entries = [{name: 'No entries found', url: '', description: '', diskUsage: 0}]
    }

    return {
        props: {entries},
    }
}

function Projects({entries}: any) {
    return <>
        <NavBar/>
        <div className="container mx-auto py-4 w-full text-white text-l">
            <ul>
                {entries.map((entry: any) => ( !entry.node.isPrivate &&
                    <li key={entry.node.name}>{entry.node.name}</li>
                ))}
            </ul>
        </div>
    </>
}

export default Projects;