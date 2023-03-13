import NavBar from "@/components/navbar";

export default function CV() {
    return <>
        <NavBar />
        <div className="container mx-auto py-4 w-full text-white text-l">
            <iframe className="w-100 vh-100" src={"cv.pdf"} />
        </div>
    </>
}