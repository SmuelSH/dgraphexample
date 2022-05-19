package mx.com.cdp.consumirws;

import androidx.appcompat.app.AppCompatActivity;

import android.os.Bundle;
import android.util.Log;
import android.view.View;
import android.widget.ArrayAdapter;
import android.widget.Button;
import android.widget.EditText;
import android.widget.ListView;
import android.widget.Toast;

import com.android.volley.AuthFailureError;
import com.android.volley.Request;
import com.android.volley.RequestQueue;
import com.android.volley.Response;
import com.android.volley.VolleyError;
import com.android.volley.toolbox.JsonArrayRequest;
import com.android.volley.toolbox.JsonObjectRequest;
import com.android.volley.toolbox.StringRequest;
import com.android.volley.toolbox.Volley;

import org.json.JSONArray;
import org.json.JSONException;
import org.json.JSONObject;
//
// import org.json.JSONObject;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class MainActivity extends AppCompatActivity {

    EditText txtUser, txtTitle, txtBody;
    Button btnEnviar;
    Button btnCharge;
    Button btnHistory;
    List<String> datos= new ArrayList<String>();

    ListView lstDatos;
    RequestQueue queque;

    String urlBuyers = "http://192.168.7.113:3000/ListBuyers";
    String urlDataCharge ="http://192.168.7.113:3000/ChargeData";
    String urlHistory ="http://192.168.7.113:3000/ListHistory";

    //datos=new ArrayList<String>();

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        txtUser = findViewById(R.id.txtUser);
        btnEnviar = findViewById(R.id.btnEnviar);
        btnCharge = findViewById(R.id.btnCharge);
        btnHistory=findViewById(R.id.btnHistory);
        queque = Volley.newRequestQueue(this);

        lstDatos = (ListView) findViewById(R.id.lstDatos);
        // ArrayAdapter<String> adapter = new ArrayAdapter<String>(this,R.layout.support_simple_spinner_dropdown_item,datos);
        //lstDatos.setAdapter(adapter);

        btnEnviar.setText("Buyers");
        btnEnviar.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View view) {
                ListBuyers();
            }
        });

        btnCharge.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View view) {
                dataCharge();
            }
        });
        txtUser.setText("820c6706"); //ejemplo
        btnHistory.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View view) {
                ListHistory();
            }
        });

    }


    private void ListBuyers(){
        //Log.e("Responsed:","fdfdfddfds");
        JsonObjectRequest request= new JsonObjectRequest(Request.Method.GET, urlBuyers, null,
                new Response.Listener<JSONObject>() {
            @Override
            public void onResponse(JSONObject response) {
                try{
                    //Log.e("Respoaa1", response.toString());
                JSONArray r=(JSONArray) response.getJSONArray("q");
                //Log.e("Len:",Integer.toString(r.length()));
                //Log.e("Len:",r.toString());
                if (r.length() > 0) {
                     Log.e("Response", response.toString());
                    datos.clear();
                    for (int i = 0; i< r.length(); i++) {
                        try {
                            JSONObject obj = r.getJSONObject(i);
                            Buyer b = new Buyer();
                            b.setIdBuyer(obj.get("idBuyer").toString());
                            b.setName(obj.get("name").toString());
                            b.setAge(obj.get("age").toString());
                            datos.add("     "+b.getIdBuyer() + "        "+b.getName()+"         "+ b.getAge());
                            ArrayAdapter<String> adapter = new ArrayAdapter<String>(getApplicationContext(),R.layout.support_simple_spinner_dropdown_item,datos);
                            lstDatos.setAdapter(adapter);
                        } catch (Exception e) {
                            e.printStackTrace();
                        }
                    }
                }
            }catch (JSONException e) {

                    e.printStackTrace();
                }
            }
        },new Response.ErrorListener() {
            @Override
            public void onErrorResponse(VolleyError error) {
                Log.e("Error", error.getMessage());
            }
        });

        queque.add(request);

    } //ListBuyers


    private void dataCharge()
    {

        StringRequest postResquest = new StringRequest(Request.Method.POST,urlDataCharge, new Response.Listener<String>() {
                    @Override
                    public void onResponse(String response) {
                        try{
                            //Log.e("Respoaa1", response.toString());
                            //JSONArray r=(JSONArray) response.getJSONArray("q");
                            //Log.e("Len:",Integer.toString(r.length()));
                            //Log.e("Len:",r.toString());
                            if (response.length() > 0) {
                                Log.e("Response", response.toString());
                                datos.clear();
                                datos.add(response.toString());
                                ArrayAdapter<String> adapter = new ArrayAdapter<String>(getApplicationContext(),R.layout.support_simple_spinner_dropdown_item,datos);
                                lstDatos.setAdapter(adapter);

                               /*
                                for (int i = 0; i< r.length(); i++) {
                                    try {
                                        JSONObject obj = r.getJSONObject(i);
                                        Buyer b = new Buyer();
                                        b.setIdBuyer(obj.get("idBuyer").toString());
                                        b.setName(obj.get("name").toString());
                                        b.setAge(obj.get("age").toString());
                                        datos.add("     "+b.getIdBuyer() + "        "+b.getName()+"         "+ b.getAge());
                                        ArrayAdapter<String> adapter = new ArrayAdapter<String>(getApplicationContext(),R.layout.support_simple_spinner_dropdown_item,datos);
                                        lstDatos.setAdapter(adapter);
                                    } catch (Exception e) {
                                        e.printStackTrace();
                                    }
                                }*/
                            }
                        }catch (Exception e) {

                            e.printStackTrace();
                        }
                    }
                },new Response.ErrorListener() {
            @Override
            public void onErrorResponse(VolleyError error) {
                Log.e("Error", error.getMessage());
            }
        });

        queque.add(postResquest);

    }//ChargeData


    private void ListHistory(){
        JSONObject paramJson = new JSONObject();
        try {
        paramJson.put("idBuyer", "820c6706");
        } catch(JSONException e){
        e.printStackTrace();
        }
        JsonObjectRequest r= new JsonObjectRequest(Request.Method.POST, urlHistory, paramJson,
                new Response.Listener<JSONObject>() {
                    @Override
                    public void onResponse(JSONObject response) {
                        try{
                            //
                            Log.e("idbuyer",txtUser.getText().toString());
                            JSONArray r=(JSONArray) response.getJSONArray("q");
                            //Log.e("Len:",Integer.toString(r.length()));
                            Log.e("Len:",r.toString());
                            if (r.length() > 0) {
                                Log.e("Response", response.toString());
                                datos.clear();

                                for (int i = 0; i< r.length(); i++) {
                                    try {
                                        JSONObject obj = r.getJSONObject(i);
                                        datos.add(obj.get("idTran").toString() + " "+obj.get("Products").toString());
                                        ArrayAdapter<String> adapter = new ArrayAdapter<String>(getApplicationContext(),R.layout.support_simple_spinner_dropdown_item,datos);
                                        lstDatos.setAdapter(adapter);
                                    } catch (Exception e) {
                                        e.printStackTrace();
                                    }
                                }
                            }
                        }catch (JSONException e) {

                            e.printStackTrace();
                        }
                    }
                },new Response.ErrorListener() {
                  @Override
                  public void onErrorResponse(VolleyError error) {
                   Log.e("Error", error.getMessage());
                 }
                 });
        queque.add(r);

    } //ListHistory

}//class



