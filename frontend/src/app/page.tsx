import { TooltipProvider } from "@/components/ui/tooltip";
import { Dashboard } from "./dashboard";

export default function Home() {
  return (
    <TooltipProvider>
      <Dashboard />
    </TooltipProvider>
  );
}
