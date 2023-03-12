import NavBar from "@/components/navbar";

export default function About(){
    return <>
        <NavBar />
        <div className="container mx-auto py-4 w-full text-white text-l">
            <h1 id="welcome-text" className="w-full">Welcome to my website!</h1>
            <p>
                Hey! I am Luuk, I am currently learning web development and I will be going back to school full-time
                starting in September.
                On this website you can find some projects I am working on and you can find more information about me.
            </p>
        </div>
    </>
}