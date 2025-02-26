import { Link } from "react-router-dom";

// componente
export function Navigation() {
  return (
    <div className="button-container">
      <Link to="/"><div>home</div></Link>

      <Link to="/info"><div>info</div></Link>
    </div>
  );
}

export default Navigation;
