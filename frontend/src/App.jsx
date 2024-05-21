import Navbar from "./components/Navbar";
import { Button } from "./components/ui/button";
import { Textarea } from "./components/ui/textarea";
import { useState } from "react";
import axios from "axios";

function App() {
    axios.defaults.baseURL = 'http://localhost:6969';

    const [currInput, setCurrInput] = useState("");
    const [responses, setResponses] = useState([]);

    async function evaluateInput(e) {
        e.preventDefault();
        console.log('here!');
        try {
            const response = await axios.post(
              '/evaluate',
              {
                'input': currInput 
              },
              {
                headers: {
                  'Content-Type': 'application/json'
                }
              }
            );

            console.log(response);

            setResponses((prevResponses) => [...prevResponses, [currInput, response.data]]);
        } catch (err) {
            console.log(err);
        };
    };

    return (
        <>
            <Navbar />  
            <section className="w-full flex items-center">
                <div className="flex flex-col w-full px-8 max-w-[1200px] mx-auto gap-6 md:gap-1 lg:gap-1 short:mt-16 tall:mt-0 pb-24">
                    <main className="flex flex-col items-center justify-center h-screen">
                        <div className="flex flex-col items-center justify-center w-full mx-auto gap-5 mt-32">
                            <Button onClick={(e) => evaluateInput(e)} className="w-full h-full">Enter Command</Button>
                            <Textarea className="w-full pb-24 resize-none" value={currInput} onChange={ (e) => setCurrInput(e.target.value) } placeholder="Feel free to input commands..." />
                        </div>
                        <div className="overflow-y-auto dark:bg-secondary w-full shadow-2xl border-2 dark:border-gray-600 rounded-xl p-4 flex flex-col gap-6 mb-8 pb-32 mt-4">
                            { responses.map((responseObject, index) => 
                                <div key={index} className="w-full">
                                    <div>
                                        {">"} {responseObject[0]}
                                    </div>
                                    <div className="">
                                        { responseObject[1].output }
                                    </div>
                                </div>
                            )}
                        </div>
                    </main>
                </div>
            </section>
        </>
    );
}

export default App
