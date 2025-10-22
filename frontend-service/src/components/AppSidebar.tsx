import { NavLink, useLocation, useNavigate } from "react-router-dom";
import { MessageSquare, User, FileText, Settings, Activity, ChevronDown, Plus } from "lucide-react";
import { Button } from "@/components/ui/button";
import { useState } from "react";

const navigationItems = [
  { title: "About You", url: "/about-you", icon: User },
  { title: "Health Records", url: "/health-records", icon: FileText },
  { title: "Activity", url: "/activity", icon: Activity },
  { title: "Settings", url: "/settings", icon: Settings },
];

// Mock chat history - in a real app, this would come from a database
const chatHistory = [
  { id: "1", title: "General Health Questions", url: "/dashboard?chat=1" },
  { id: "2", title: "Nutrition Advice", url: "/dashboard?chat=2" },
  { id: "3", title: "Exercise Routine", url: "/dashboard?chat=3" },
];

export function AppSidebar() {
  const location = useLocation();
  const navigate = useNavigate();
  const [chatsExpanded, setChatsExpanded] = useState(true);
  
  const isActive = (path: string) => location.pathname === path;
  const isChatActive = location.pathname === "/dashboard";
  
  const handleNewChat = () => {
    navigate("/dashboard");
    // In a real app, this would create a new chat session
  };
  
  return (
    <aside className="sticky top-0 w-64 h-screen bg-sidebar border-r border-sidebar-border flex flex-col">
      {/* Logo Section */}
      <div className="p-6 border-b border-sidebar-border">
        <h1 className="text-2xl font-bold text-sidebar-foreground">Health Harbor</h1>
        <p className="text-xs text-sidebar-foreground/70 mt-1">Your Personal Health Advisor</p>
      </div>
      
      {/* Navigation */}
      <nav className="flex-1 p-4 space-y-2">
        {/* Chats Section */}
        <div className="space-y-1">
          <div className="flex items-center gap-2">
            <Button
              variant="nav"
              size="nav"
              onClick={handleNewChat}
              className={`flex-1 ${isChatActive ? "bg-primary/10 font-medium" : ""}`}
            >
              <MessageSquare className="mr-3 h-5 w-5" />
              Chats
            </Button>
            <Button
              variant="ghost"
              size="icon"
              onClick={() => setChatsExpanded(!chatsExpanded)}
              className="h-10 w-10 hover:bg-sidebar-accent"
            >
              <ChevronDown className={`h-4 w-4 transition-transform ${chatsExpanded ? "" : "-rotate-90"}`} />
            </Button>
          </div>
          
          {/* Chat History */}
          {chatsExpanded && (
            <div className="ml-4 space-y-1">
              {chatHistory.map((chat) => (
                <NavLink key={chat.id} to={chat.url}>
                  <Button
                    variant="ghost"
                    size="sm"
                    className="w-full justify-start text-xs text-sidebar-foreground/70 hover:bg-sidebar-foreground/10 hover:text-sidebar-foreground"
                  >
                    {chat.title}
                  </Button>
                </NavLink>
              ))}
            </div>
          )}
        </div>
        
        {/* Other Navigation Items */}
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
