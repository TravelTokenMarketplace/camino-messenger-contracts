import { useParams } from "react-router-dom";
import { Card } from "../components/Card";

export function AccountWorkspace() {
  const { address } = useParams();
  return <Card title="Account Workspace">{address}</Card>;
}
