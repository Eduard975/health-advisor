import { NavLink, useLocation } from "react-router-dom";
import { Home, User, FileText, Settings, Activity } from "lucide-react";
import { Button } from "@/components/ui/button";

const navigationItems = [
  { title: "Dashboard", url: "/dashboard", icon: Home },
  { title: "About You", url: "/about-you", icon: User },
  { title: "Health Records", url: "/health-records", icon: FileText },
  { title: "Activity", url: "/activity", icon: Activity },
  { title: "Settings", url: "/settings", icon: Settings },
];

export function AppSidebar() {
  const location = useLocation();
  
  const isActive = (path: string) => location.pathname === path;
  
  return (
    <aside className="sticky top-0 w-64 h-screen bg-sidebar border-r border-sidebar-border flex flex-col">
      {/* Logo Section */}
      <div className="p-6 border-b border-sidebar-border">
        <h1 className="text-2xl font-bold text-sidebar-foreground">Health Harbor</h1>
        <p className="text-xs text-sidebar-foreground/70 mt-1">Your Personal Health Advisor</p>
      </div>
      
      {/* Navigation */}
      <nav className="flex-1 p-4 space-y-2">
        {navigationItems.map((item) => {
          const Icon = item.icon;
          const active = isActive(item.url);
          
          return (
            <NavLink key={item.url} to={item.url}>
              <Button
                variant="nav"
                size="nav"
                className={active ? "bg-sidebar-accent font-medium" : ""}
              >
                <Icon className="mr-3 h-5 w-5" />
                {item.title}
              </Button>
            </NavLink>
          );
        })}
      </nav>
      
      {/* Footer */}
      <div className="p-4 border-t border-sidebar-border">
        <p className="text-xs text-sidebar-foreground/60 text-center">
          © 2025 Health Harbor
        </p>
      </div>
    </aside>
  );
}
