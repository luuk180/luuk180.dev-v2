import Link from "next/link";
import {useRouter} from "next/router";
import styles from '@/styles/navbar.module.css';

export default function NavBar(){
    const router = useRouter();
    return <div>
        <nav className="navbar navbar-expand-sm bg-dark navbar-dark">
            <div className="container-fluid">
                <div className={styles.logo}>
                    <a className="navbar-brand">
                        luuk180
                    </a>
                </div>
                <button className="navbar-toggler" type="button" data-bs-toggle="collapse"
                        data-bs-target="#collapsibleNavbar">
                    <span className="navbar-toggler-icon"></span>
                </button>
                <div className="collapse navbar-collapse" id="collapsibleNavbar">
                    <ul className="navbar-nav">
                        <li>
                            {
                                router.route === "/" ?
                                    <Link className="nav-link px-3 text-secondary" href="/">Home</Link>:
                                    <Link className="nav-link px-3 text-white" href="/">Home</Link>
                            }
                        </li>
                        <li>
                            {
                                router.route === "/projects" ?
                                    <Link className="nav-link px-3 text-secondary" href="/projects">Projects</Link>:
                                    <Link className="nav-link px-3 text-white" href="/projects">Projects</Link>
                            }
                        </li>
                        <li>
                            {
                                router.route === "/cv" ?
                                    <Link className="nav-link px-3 text-secondary" href="/cv">CV</Link>:
                                    <Link className="nav-link px-3 text-white" href="/cv">CV</Link>
                            }
                        </li>
                        <li>
                            {
                                router.route === "/about" ?
                                    <Link className="nav-link px-3 text-secondary" href="/about">About</Link>:
                                    <Link className="nav-link px-3 text-white" href="/about">About</Link>
                            }
                        </li>
                    </ul>
                </div>
            </div>
        </nav>
    </div>
}