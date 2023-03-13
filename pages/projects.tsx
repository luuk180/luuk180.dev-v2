import NavBar from "@/components/navbar";
import {GetServerSideProps} from "next";

export const getServerSideProps: GetServerSideProps = async ({res}) => {
    res.setHeader('Cache-Control', 'public, s-maxage=300, stale-while-revalidate=600');
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
                         pushedAt
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
        <div className="container mx-auto py-4 w-full text-l">
            <div className="card-group">
                {entries.map((entry: any) => ( !entry.node.isPrivate &&
                    <div key={entry.node.pushedAt} className="card">
                        <div className="card-header">
                          <a className="btn card-title" href={entry.node.url}>{entry.node.name}</a>
                        </div>
                        <p className="card-text">{entry.node.description}</p>
                        <p className="card-footer">{entry.node.diskUsage} kB</p>
                    </div>
                )).reverse()}
            </div>
        </div>
    </>
}

export default Projects;