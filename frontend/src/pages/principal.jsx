// src/pages/Principal.jsx
import Header from "../components/Header";

import InputConsole from "../components/InputConsole";
import OutputConsole from "../components/OutputConsole";
import FileUpload from "../components/FileUpload";

import ExecuteButton from "../components/ExecuteButton";
import ClearButton from "../components/ClearButton";

import Footer from "../components/Footer";

import { Navigation } from "../components/Navigation";


export function Principal() {
  return (
    <div>
      <Header />
      <div>
        <FileUpload />
        <ExecuteButton />
        <ClearButton />
        
      </div>
      <main>
        <section>
          <h3>Entrada:</h3>
          <InputConsole />
        </section>
        <section>
          <h3>Salida:</h3>
          <OutputConsole />
        </section>
      </main>
      <Footer />
    </div>
  );
}

export default Principal;
