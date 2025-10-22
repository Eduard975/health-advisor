import { AppSidebar } from "@/components/AppSidebar";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Label } from "@/components/ui/label";
import { Switch } from "@/components/ui/switch";
import { Input } from "@/components/ui/input";
import { toast } from "@/hooks/use-toast";
import { Settings as SettingsIcon, Bell, Lock, User, Smartphone } from "lucide-react";
import defaultImage from "@/assets/default.jpg";

const Settings = () => {
  const handleSaveSettings = () => {
    toast({
      title: "Settings Saved",
      description: "Your preferences have been updated successfully.",
    });
  };

  return (
    <div className="flex min-h-screen w-full bg-background">
      <AppSidebar />
      
      <main className="flex-1 overflow-auto">
        {/* Header */}
        <header className="border-b border-border bg-card">
          <div className="p-6">
            <div className="flex items-center gap-4">
              <div className="h-12 w-12 rounded-full bg-primary/10 flex items-center justify-center">
                <SettingsIcon className="h-6 w-6 text-primary" />
              </div>
              <div>
                <h1 className="text-3xl font-bold text-foreground">Settings</h1>
                <p className="text-muted-foreground mt-1">
                  Manage your account preferences and notifications
                </p>
              </div>
            </div>
          </div>
        </header>

        <div className="p-6 max-w-4xl mx-auto space-y-6">
          {/* Profile Settings */}
          <Card>
            <CardHeader>
              <div className="flex items-center gap-3">
                <User className="h-6 w-6 text-primary" />
                <div>
                  <CardTitle>Profile Settings</CardTitle>
                  <CardDescription>Update your personal information and photo</CardDescription>
                </div>
              </div>
            </CardHeader>
            <CardContent className="space-y-4">
              <div className="flex items-center gap-6">
                <img
                  src={defaultImage}
                  alt="Profile"
                  className="h-20 w-20 rounded-full object-cover ring-4 ring-primary"
                />
                <Button variant="secondary">Change Photo</Button>
              </div>
              
              <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                <div className="space-y-2">
                  <Label htmlFor="email">Email Address</Label>
                  <Input
                    id="email"
                    type="email"
                    defaultValue="user@example.com"
                    className="focus-fade"
                  />
                </div>
                
                <div className="space-y-2">
                  <Label htmlFor="phone">Phone Number</Label>
                  <Input
                    id="phone"
                    type="tel"
                    defaultValue="+1 (555) 123-4567"
                    className="focus-fade"
                  />
                </div>
              </div>
            </CardContent>
          </Card>

          {/* Notification Settings */}
          <Card>
            <CardHeader>
              <div className="flex items-center gap-3">
                <Bell className="h-6 w-6 text-accent" />
                <div>
                  <CardTitle>Notifications</CardTitle>
                  <CardDescription>Choose how you want to be notified</CardDescription>
                </div>
              </div>
            </CardHeader>
            <CardContent className="space-y-6">
              <div className="flex items-center justify-between">
                <div className="space-y-0.5">
                  <Label htmlFor="email-notifications" className="font-medium">
                    Email Notifications
                  </Label>
                  <p className="text-sm text-muted-foreground">
                    Receive appointment reminders via email
                  </p>
                </div>
                <Switch id="email-notifications" defaultChecked />
              </div>

              <div className="flex items-center justify-between">
                <div className="space-y-0.5">
                  <Label htmlFor="push-notifications" className="font-medium">
                    Push Notifications
                  </Label>
                  <p className="text-sm text-muted-foreground">
                    Get alerts about health insights
                  </p>
                </div>
                <Switch id="push-notifications" defaultChecked />
              </div>

              <div className="flex items-center justify-between">
                <div className="space-y-0.5">
                  <Label htmlFor="activity-reminders" className="font-medium">
                    Activity Reminders
                  </Label>
                  <p className="text-sm text-muted-foreground">
                    Daily reminders to log your health metrics
                  </p>
                </div>
                <Switch id="activity-reminders" />
              </div>

              <div className="flex items-center justify-between">
                <div className="space-y-0.5">
                  <Label htmlFor="medication-alerts" className="font-medium">
                    Medication Alerts
                  </Label>
                  <p className="text-sm text-muted-foreground">
                    Reminders to take your medications
                  </p>
                </div>
                <Switch id="medication-alerts" defaultChecked />
              </div>
            </CardContent>
          </Card>

          {/* Privacy & Security */}
          <Card>
            <CardHeader>
              <div className="flex items-center gap-3">
                <Lock className="h-6 w-6 text-secondary" />
                <div>
                  <CardTitle>Privacy & Security</CardTitle>
                  <CardDescription>Manage your account security settings</CardDescription>
                </div>
              </div>
            </CardHeader>
            <CardContent className="space-y-4">
              <div className="flex items-center justify-between">
                <div className="space-y-0.5">
                  <Label htmlFor="two-factor" className="font-medium">
                    Two-Factor Authentication
                  </Label>
                  <p className="text-sm text-muted-foreground">
                    Add an extra layer of security to your account
                  </p>
                </div>
                <Switch id="two-factor" />
              </div>

              <div className="pt-4 border-t">
                <Button variant="secondary" className="w-full md:w-auto">
                  Change Password
                </Button>
              </div>
            </CardContent>
          </Card>

          {/* Connected Devices */}
          <Card>
            <CardHeader>
              <div className="flex items-center gap-3">
                <Smartphone className="h-6 w-6 text-primary" />
                <div>
                  <CardTitle>Connected Devices</CardTitle>
                  <CardDescription>Manage devices linked to your account</CardDescription>
                </div>
              </div>
            </CardHeader>
            <CardContent>
              <div className="space-y-4">
                <div className="flex items-center justify-between p-4 bg-muted rounded-lg">
                  <div>
                    <p className="font-medium text-foreground">iPhone 14 Pro</p>
                    <p className="text-sm text-muted-foreground">Last active: 2 minutes ago</p>
                  </div>
                  <Button variant="ghost" size="sm">Remove</Button>
                </div>
                
                <div className="flex items-center justify-between p-4 bg-muted rounded-lg">
                  <div>
                    <p className="font-medium text-foreground">MacBook Pro</p>
                    <p className="text-sm text-muted-foreground">Last active: 1 hour ago</p>
                  </div>
                  <Button variant="ghost" size="sm">Remove</Button>
                </div>
              </div>
            </CardContent>
          </Card>

          {/* Save Button */}
          <div className="flex justify-end gap-3 pt-4">
            <Button variant="secondary">Cancel</Button>
            <Button variant="accent" size="lg" onClick={handleSaveSettings}>
              Save Changes
            </Button>
          </div>
        </div>
      </main>
    </div>
  );
};

export default Settings;
