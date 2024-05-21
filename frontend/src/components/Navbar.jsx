import React from "react";
import { ModeToggle } from "./ModeToggle";
import { Separator } from "@/components/ui/separator";

function Navbar() {
    return (
        <header className="fixed top-0 left-0 right-0 flex flex-col items-center backdrop-filter backdrop-blur-lg bg-opacity-30 px-4 md:px-6 lg:px-8 z-20">
            <nav className="flex justify-between gap-14 w-full max-w-[1200px] py-4 px-4 items-center">
                <div className="flex gap-2 cursor-pointer items-center z-30">
                    <h3 className="text-lg hidden lg:flex font-semibold tracking-tight">
                        While Language 
                    </h3>
                </div>
                <div className="flex gap-4">
                    <ModeToggle />
                </div>
            </nav>
            <Separator className="w-screen"/>
        </header>
    );
};

export default Navbar;
