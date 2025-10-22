import { AppSidebar } from "@/components/AppSidebar";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { FileText, Download, Plus, Calendar } from "lucide-react";
import defaultImage from "@/assets/default.jpg";

// Mock data for health records
const healthRecords = [
  {
    id: 1,
    title: "Annual Physical Examination",
    date: "2025-01-15",
    doctor: "Dr. Sarah Johnson",
    type: "Checkup",
    status: "completed",
  },
  {
    id: 2,
    title: "Blood Test Results",
    date: "2025-01-10",
    doctor: "Dr. Michael Chen",
    type: "Lab Work",
    status: "completed",
  },
  {
    id: 3,
    title: "Cardiology Consultation",
    date: "2024-12-20",
    doctor: "Dr. Emily Rodriguez",
    type: "Specialist",
    status: "completed",
  },
  {
    id: 4,
    title: "Vaccination Record - Flu Shot",
    date: "2024-11-05",
    doctor: "Dr. Sarah Johnson",
    type: "Immunization",
    status: "completed",
  },
];

const upcomingAppointments = [
  {
    id: 1,
    title: "Follow-up Consultation",
    date: "2025-02-10",
    time: "10:30 AM",
    doctor: "Dr. Sarah Johnson",
  },
  {
    id: 2,
    title: "Dental Checkup",
    date: "2025-02-18",
    time: "2:00 PM",
    doctor: "Dr. James Wilson",
  },
];

const HealthRecords = () => {
  return (
    <div className="flex min-h-screen w-full bg-background">
      <AppSidebar />
      
      <main className="flex-1 overflow-auto">
        {/* Header */}
        <header className="border-b border-border bg-card">
          <div className="p-6">
            <div className="flex items-center justify-between">
              <div>
                <h1 className="text-3xl font-bold text-foreground">Health Records</h1>
                <p className="text-muted-foreground mt-1">
                  View and manage your medical documents and appointments
                </p>
              </div>
              <div className="flex gap-3">
                <Button variant="secondary" size="lg">
                  <Calendar className="mr-2 h-4 w-4" />
                  Add Appointment
                </Button>
                <Button variant="accent" size="lg">
                  <Plus className="mr-2 h-4 w-4" />
                  Upload Record
                </Button>
              </div>
            </div>
          </div>
        </header>

        <div className="p-6 max-w-7xl mx-auto space-y-6">
          {/* Upcoming Appointments */}
          <Card>
            <CardHeader>
              <div className="flex items-center justify-between">
                <div>
                  <CardTitle className="text-2xl">Upcoming Appointments</CardTitle>
                  <CardDescription>Your scheduled medical appointments</CardDescription>
                </div>
                <Calendar className="h-8 w-8 text-primary" />
              </div>
            </CardHeader>
            <CardContent>
              <div className="space-y-4">
                {upcomingAppointments.map((appointment) => (
                  <div
                    key={appointment.id}
                    className="flex items-center justify-between p-4 bg-muted rounded-lg hover:bg-muted/70 transition-colors"
                  >
                    <div className="flex items-center gap-4">
                      <img
                        src={defaultImage}
                        alt={appointment.doctor}
                        className="h-12 w-12 rounded-full object-cover ring-2 ring-primary"
                      />
                      <div>
                        <h3 className="font-medium text-foreground">{appointment.title}</h3>
                        <p className="text-sm text-muted-foreground">{appointment.doctor}</p>
                      </div>
                    </div>
                    <div className="text-right">
                      <p className="font-medium text-foreground">{appointment.date}</p>
                      <p className="text-sm text-muted-foreground">{appointment.time}</p>
                    </div>
                  </div>
                ))}
              </div>
            </CardContent>
          </Card>

          {/* Medical Records */}
          <Card>
            <CardHeader>
              <div className="flex items-center justify-between">
                <div>
                  <CardTitle className="text-2xl">Medical Records</CardTitle>
                  <CardDescription>Your complete medical history and documents</CardDescription>
                </div>
                <FileText className="h-8 w-8 text-secondary" />
              </div>
            </CardHeader>
            <CardContent>
              <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                {healthRecords.map((record) => (
                  <div
                    key={record.id}
                    className="border border-border rounded-lg p-4 hover:border-primary transition-colors"
                  >
                    <div className="flex items-start justify-between mb-3">
                      <div className="flex-1">
                        <h3 className="font-medium text-foreground mb-1">{record.title}</h3>
                        <p className="text-sm text-muted-foreground">{record.doctor}</p>
                      </div>
                      <Badge variant="secondary">{record.type}</Badge>
                    </div>
                    
                    <div className="flex items-center justify-between">
                      <p className="text-sm text-muted-foreground">{record.date}</p>
                      <Button variant="ghost" size="sm">
                        <Download className="h-4 w-4" />
                      </Button>
                    </div>
                  </div>
                ))}
              </div>
            </CardContent>
          </Card>

          {/* Quick Stats */}
          <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
            <Card>
              <CardHeader>
                <CardDescription>Total Records</CardDescription>
                <CardTitle className="text-4xl text-primary">24</CardTitle>
              </CardHeader>
            </Card>
            
            <Card>
              <CardHeader>
                <CardDescription>Upcoming Visits</CardDescription>
                <CardTitle className="text-4xl text-secondary">2</CardTitle>
              </CardHeader>
            </Card>
            
            <Card>
              <CardHeader>
                <CardDescription>Last Checkup</CardDescription>
                <CardTitle className="text-4xl text-accent">Jan 15</CardTitle>
              </CardHeader>
            </Card>
          </div>
        </div>
      </main>
    </div>
  );
};

export default HealthRecords;
